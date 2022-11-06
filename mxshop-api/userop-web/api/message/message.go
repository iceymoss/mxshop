package message

import (
	"context"
	"net/http"

	"mxshop-api/user-web/api"
	"mxshop-api/userop-web/forms"
	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/models"
	"mxshop-api/userop-web/proto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	claimsInfo, _ := ctx.Get("claims")
	claims := claimsInfo.(*models.CustomClaims)
	request := proto.MessageRequest{}
	if claims.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	Rsp, err := global.MessageSrvClient.GetMessageList(context.Background(), &request)
	if err != nil {
		zap.S().Info("")
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}

	ReMap := make(map[string]interface{})
	ReMap["total"] = Rsp.Total
	messageList := make([]interface{}, 0)
	for _, message := range Rsp.Data {
		messageInfo := make(map[string]interface{})
		messageInfo["id"] = message.Id
		messageInfo["user_id"] = message.UserId
		messageInfo["message_type"] = message.MessageType
		messageInfo["subject"] = message.Subject
		messageInfo["message"] = message.Message
		messageInfo["file"] = message.File
		messageList = append(messageList, messageInfo)
	}
	ReMap["data"] = messageList

	ctx.JSON(http.StatusOK, ReMap)
}

func New(ctx *gin.Context) {
	var message forms.MessageForm
	if err := ctx.ShouldBindJSON(&message); err != nil {
		zap.S().Info("获取表单失败")
		api.HandleValidatorErr(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")

	Rsp, err := global.MessageSrvClient.CreateMessage(context.Background(), &proto.MessageRequest{
		UserId:      int32(userId.(uint)),
		MessageType: message.MessageType,
		Subject:     message.Subject,
		Message:     message.Message,
		File:        message.File,
	})
	if err != nil {
		zap.S().Info("新建留言失败", err)
		api.HandleGrpcErrToHttp(err, ctx)
		return
	}
	request := make(map[string]interface{})
	request["id"] = Rsp.Id

	ctx.JSON(http.StatusOK, request)
}
