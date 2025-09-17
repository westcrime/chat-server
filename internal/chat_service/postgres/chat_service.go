package postgres

import (
	"context"

	chatrepository "github.com/westcrime/chat-server/internal/chat_repository"
	chatservice "github.com/westcrime/chat-server/internal/chat_service"
	"github.com/westcrime/chat-server/internal/model"
)

type chatService struct {
	chatRepository chatrepository.ChatRepository
}

func NewChatService(chatRepository chatrepository.ChatRepository) chatservice.ChatService {
	return &chatService{chatRepository: chatRepository}
}

func (cs *chatService) Create(ctx context.Context, createChat *model.CreateChat) (error, int64) {
	return cs.chatRepository.CreateChat(ctx, createChat)
}

func (cs *chatService) Delete(ctx context.Context, id int64) error {
	return cs.chatRepository.DeleteChat(ctx, id)
}

func (cs *chatService) SendMessage(ctx context.Context, message *model.Message) error {
	return cs.chatRepository.SendMessage(ctx, message)
}
