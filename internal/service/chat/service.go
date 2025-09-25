package chat

import (
	"github.com/westcrime/chat-server/internal/repository"
	"github.com/westcrime/chat-server/internal/service"
)

type chatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) service.ChatService {
	return &chatService{chatRepository: chatRepository}
}
