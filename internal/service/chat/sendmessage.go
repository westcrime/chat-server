package chat

import (
	"context"

	"github.com/westcrime/chat-server/internal/model"
)

func (cs *chatService) SendMessage(ctx context.Context, message *model.Message) error {
	return cs.chatRepository.SendMessage(ctx, message)
}
