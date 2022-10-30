package initialize

import (
	"context"
	"fmt"
	"log"
	"os"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

//InitEs 初始化es服务
func InitEs() {
	EsConfig := global.ServerConfig.EsConfig
	url := fmt.Sprintf("http://%s:%d/", EsConfig.Host, EsConfig.Port)
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)

	var err error
	global.EsClient, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		zap.S().Info("初始化EsClient失败:", err)
	}

	//新建index和mapping
	//先查询index是否存在
	exists, err := global.EsClient.IndexExists(model.EsGoods{}.GetIndexName()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//新建索引和mapping
	if !exists {
		_, err = global.EsClient.CreateIndex(model.EsGoods{}.GetIndexName()).BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

}
