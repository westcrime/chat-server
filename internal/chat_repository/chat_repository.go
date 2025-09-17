package chatrepository

import (
	"context"

	"github.com/westcrime/chat-server/internal/model"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, createChat *model.CreateChat) (error, int64)
	SendMessage(ctx context.Context, SendMessage *model.Message) error
	DeleteChat(ctx context.Context, user_id int64) error
}
