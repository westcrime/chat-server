package grpcserver

import (
	"context"
	"log"

	chatservice "github.com/westcrime/chat-server/internal/chat_service"
	"github.com/westcrime/chat-server/internal/converter"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	desc.UnimplementedChatV1Server
	chatService chatservice.ChatService
}

func NewServer(usp chatservice.ChatService) *Server {
	return &Server{chatService: usp}
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat id: %d", req.GetId())

	err := s.chatService.Delete(ctx, req.GetId())

	return &emptypb.Empty{}, err
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Chat: %v", req.Info)
	err, id := s.chatService.Create(ctx, converter.ToCreateChatModelFromCreateRequestProto(req))

	return &desc.CreateResponse{
		Id: id}, err
}

func (s *Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message: %v", req)
	err := s.chatService.SendMessage(ctx, converter.ToSendMessageModelFromUSendMessageRequestProto(req))

	return &emptypb.Empty{}, err
}
