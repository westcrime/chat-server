package grpcserver

import (
	"context"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *ChatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat id: %d", req.GetId())

	err := s.chatService.Delete(ctx, req.GetId())

	return &emptypb.Empty{}, err
}
