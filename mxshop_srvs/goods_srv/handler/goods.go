package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//ModelToResponse model转response
func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

//GoodsList 商品列表
func (g *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	//var goodsListResponse proto.GoodsListResponse
	//先根据条件去es接口拼接成es语法，去es中查询符合条件的数据，然后返回对应数据的id
	//拿着es返回的商品id去mysql获取到商品

	goodsListResponse := &proto.GoodsListResponse{}

	//初始化筛选器
	q := elastic.NewBoolQuery()
	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选
	//创建一个临时筛选条件的DB
	var localDB = global.DB.Model(model.Goods{})

	if req.KeyWords != "" {
		q = q.Must(elastic.NewMultiMatchQuery(req.KeyWords, "name", "goods_brief")) //在name和goods_brief查询
	}
	if req.IsNew {
		//不参与算分，使用filter
		q = q.Filter(elastic.NewTermsQuery("is_new", req.IsNew))
	}
	if req.IsHot {
		q = q.Filter(elastic.NewTermsQuery("is_hot", req.IsHot))
	}
	if req.PriceMin > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Gte(req.PriceMin))
	}
	if req.PriceMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Lte(req.PriceMax))
	}
	if req.Brand > 0 {
		q = q.Filter(elastic.NewTermsQuery("brand_id", req.Brand))
	}

	//根据类目级别查找
	var SubQuery string
	categoryIds := make([]interface{}, 0)
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		//根据类目级别返回查询子分类
		if category.Level == 1 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id in (SELECT id FROM category WHERE parent_category_id = %d)", req.TopCategory)
		} else if category.Level == 2 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id = %d", req.TopCategory)
		} else if category.Level == 3 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", req.TopCategory)
		}

		//这里需要将类目级别相关信息获取出相应的数据，用于在es查询相应商品信息
		type result struct {
			ID int32
		}
		var Results []result

		//将SubQuery语句获取到的Category映射到Result
		//获取对应分类的子分类id
		global.DB.Model(model.Category{}).Raw(SubQuery).Scan(&Results)
		for _, p := range Results {
			categoryIds = append(categoryIds, p.ID)
		}

		fmt.Println("categoryID:", categoryIds)

		//将目录级别分类条件放入q
		q = q.Filter(elastic.NewTermsQuery("category_id", categoryIds...))
	}

	//分页：From().Size()
	if req.Pages == 0 {
		req.Pages = 1
	}
	switch {
	case req.PagePerNums > 100:
		req.PagePerNums = 100
	case req.PagePerNums <= 0:
		req.PagePerNums = 10
	}

	fmt.Println("es搜索:", q)

	//es根据q条件查询
	re, err := global.EsClient.Search().Index(model.EsGoods{}.GetIndexName()).Query(q).From(int(req.Pages)).Size(int(req.PagePerNums)).Do(context.Background())
	if err != nil {
		return nil, err
	}

	goodsListResponse.Total = int32(re.Hits.TotalHits.Value)
	goodIDs := make([]int32, 0)

	fmt.Println("共计：", goodsListResponse.Total)
	for _, v := range re.Hits.Hits {
		goods := model.EsGoods{}
		json.Unmarshal(v.Source, &goods)
		goodIDs = append(goodIDs, goods.ID)
	}

	fmt.Println("goodsID:", goodIDs)

	var Goods []model.Goods

	//查询
	Result := localDB.Preload("Category").Preload("Brands").Find(&Goods, goodIDs)
	if Result.Error != nil {
		return nil, Result.Error
	}

	for _, good := range Goods {
		GoodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &GoodsInfoResponse)
	}

	return goodsListResponse, nil
}

//BatchGetGoods 查询多个商品信息
func (g *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	//用户提交订单有多个商品，需要批量查询商品的信息

	var GoodsListResponse proto.GoodsListResponse
	var goods []model.Goods

	//调用where并不会真正执行sql 只是用来生成sql的 当调用find， first才会去执行sql，
	result := global.DB.Where(req.Id).Find(&goods)
	for _, good := range goods {
		GoodsInfoResponse := ModelToResponse(good)
		GoodsListResponse.Data = append(GoodsListResponse.Data, &GoodsInfoResponse)
	}
	GoodsListResponse.Total = int32(result.RowsAffected)
	return &GoodsListResponse, nil
}

//CreateGoods 添加商品
func (g *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	var goods model.Goods

	goods.Brands = brand
	goods.BrandsID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID

	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale
	goods.ShipFree = req.ShipFree

	goods.ID = req.Id
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.GoodsBrief = req.GoodsBrief
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.DescImages = req.DescImages
	goods.Images = req.Images

	//这里需要利用钩子和事务将mysql和es数据保持一致性
	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		//业务回滚
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &proto.GoodsInfoResponse{
		Id: goods.ID,
	}, nil
}

//DeleteGoods 删除商品
func (g *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*empty.Empty, error) {
	tx := global.DB.Begin()
	if result := tx.Delete(&model.Goods{BaseModel: model.BaseModel{ID: req.Id}}); result.RowsAffected == 0 {
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}

//UpdateGoods 更新商品信息
func (g *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*empty.Empty, error) {
	var goods model.Goods
	if result := global.DB.First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	goods.Brands = brand
	goods.BrandsID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID

	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale
	goods.ShipFree = req.ShipFree

	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.GoodsBrief = req.GoodsBrief
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.DescImages = req.DescImages
	goods.Images = req.Images

	tx := global.DB.Begin()
	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &emptypb.Empty{}, nil
}

//GetGoodsDetail 根据id查询商品
func (g *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var goods model.Goods

	//多表关联查询需要进行Preload预加载
	if result := global.DB.Preload("Category").Preload("Brands").First(&goods, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	GoodsInfoRes := ModelToResponse(goods)
	return &GoodsInfoRes, nil
}
