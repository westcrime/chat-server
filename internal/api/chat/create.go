package grpcserver

import (
	"context"
	"log"

	"github.com/westcrime/chat-server/internal/api/chat/converter"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
)

func (s *ChatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Chat: %v", req.Info)
	err, id := s.chatService.Create(ctx, converter.ToCreateChatModelFromCreateRequestProto(req))

	return &desc.CreateResponse{
		Id: id}, err
}
