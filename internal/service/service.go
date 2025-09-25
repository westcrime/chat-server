package service

import (
	"context"

	"github.com/westcrime/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, createChat *model.CreateChat) (error, int64)
	Delete(ctx context.Context, chat_id int64) error
	SendMessage(ctx context.Context, message *model.Message) error
}
