package handler

import (
	"context"

	"mxshop_srvs/userop_srv/global"
	"mxshop_srvs/userop_srv/model"
	"mxshop_srvs/userop_srv/proto"
)

//GetMessageList 获取留言
func (u *UserOpServer) GetMessageList(ctx context.Context, req *proto.MessageRequest) (*proto.MessageListResponse, error) {
	var rsp proto.MessageListResponse
	var messages []model.LeavingMessages
	var messageList []*proto.MessageResponse

	result := global.DB.Where(&model.LeavingMessages{User: req.UserId}).Find(&messages)
	rsp.Total = int32(result.RowsAffected)

	for _, message := range messages {
		messageList = append(messageList, &proto.MessageResponse{
			Id:          message.ID,
			UserId:      message.User,
			MessageType: message.MessageType,
			Subject:     message.Subject,
			Message:     message.Message,
			File:        message.File,
		})
	}

	rsp.Data = messageList
	return &rsp, nil
}

//CreateMessage 新建留言
func (u *UserOpServer) CreateMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	var message model.LeavingMessages

	message.User = req.UserId
	message.MessageType = req.MessageType
	message.Subject = req.Subject
	message.Message = req.Message
	message.File = req.File

	global.DB.Save(&message)

	return &proto.MessageResponse{Id: message.ID}, nil
}
