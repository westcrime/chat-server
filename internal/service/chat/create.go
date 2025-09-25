package chat

import (
	"context"

	"github.com/westcrime/chat-server/internal/model"
)

func (cs *chatService) Create(ctx context.Context, createChat *model.CreateChat) (error, int64) {
	return cs.chatRepository.CreateChat(ctx, createChat)
}
