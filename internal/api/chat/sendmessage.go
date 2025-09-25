package grpcserver

import (
	"context"
	"log"

	"github.com/westcrime/chat-server/internal/api/chat/converter"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ChatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message: %v", req)
	err := s.chatService.SendMessage(ctx, converter.ToSendMessageModelFromUSendMessageRequestProto(req))

	return &emptypb.Empty{}, err
}
