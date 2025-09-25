package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jackc/pgx/v5/pgxpool"
	urp "github.com/westcrime/chat-server/internal/repository/chat"
	usp "github.com/westcrime/chat-server/internal/service/chat"
	"github.com/westcrime/chat-server/internal/config"
	api "github.com/westcrime/chat-server/internal/api/chat"
	desc "github.com/westcrime/chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}
	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	ur := urp.NewChatRepository(pool)
	us := usp.NewChatService(ur)

	desc.RegisterChatV1Server(s, api.NewChatServer(us))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
