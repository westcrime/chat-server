package chat

import (
	"context"
)

func (cs *chatService) Delete(ctx context.Context, id int64) error {
	return cs.chatRepository.DeleteChat(ctx, id)
}
