package converter

import (
	"github.com/westcrime/chat-server/internal/model"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
)

func ToCreateChatModelFromCreateRequestProto(req *desc.CreateRequest) *model.CreateChat {
	return &model.CreateChat{
		Name:      req.Info.Name,
		Usernames: req.Info.Usernames,
	}
}

func ToSendMessageModelFromUSendMessageRequestProto(req *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		From:      req.From,
		Text:      req.Text,
		Timestamp: req.Timestamp.AsTime(),
		ChatId:    req.ChatId,
	}
}
