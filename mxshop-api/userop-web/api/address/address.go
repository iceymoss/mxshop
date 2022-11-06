package address

import (
	"context"
	"net/http"
	"strconv"

	"mxshop-api/userop-web/api"
	"mxshop-api/userop-web/forms"
	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/models"
	"mxshop-api/userop-web/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(ctx *gin.Context) {
	request := &proto.GetAddressRequest{}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	if currentUser.AuthorityId != 2 {
		userId, _ := ctx.Get("userId")
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.AddressSrvClient.GetAddressList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("获取地址列表失败")
		api.HandleValidatorErr(ctx, err)
		return
	}

	reMap := map[string]interface{}{
		"total": rsp.Total,
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		Map := make(map[string]interface{})
		Map["id"] = value.Id
		Map["user_id"] = value.UserId
		Map["province"] = value.Province
		Map["city"] = value.City
		Map["district"] = value.District
		Map["address"] = value.Address
		Map["signer_name"] = value.SignerName
		Map["signer_mobile"] = value.SignerMobile

		result = append(result, Map)
	}

	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)

}

func New(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		api.HandleValidatorErr(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")

	rsp, err := global.AddressSrvClient.CreateAddress(context.Background(), &proto.GetAddressRequest{
		UserId:       int32(userId.(uint)),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})

	if err != nil {
		zap.S().Errorw("新建地址失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id

	ctx.JSON(http.StatusOK, request)

}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")

	_, err = global.AddressSrvClient.DeleteAddress(context.Background(), &proto.GetAddressRequest{
		Id:     int32(i),
		UserId: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("删除地址失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, "")

}

func Update(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		api.HandleValidatorErr(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	userId, _ := ctx.Get("userId")

	_, err = global.AddressSrvClient.UpdateAddress(context.Background(), &proto.GetAddressRequest{
		Id:           int32(i),
		UserId:       int32(userId.(uint)),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("更新地址失败")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}
	request := make(map[string]interface{})
	ctx.JSON(http.StatusOK, request)
}
