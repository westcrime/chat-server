package grpcserver

import (
	"github.com/westcrime/chat-server/internal/service"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
)

type ChatServer struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewChatServer(usp service.ChatService) *ChatServer {
	return &ChatServer{chatService: usp}
}
