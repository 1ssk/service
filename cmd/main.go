package main

import (
	"context"
	"fmt"
	"github.com/test/integnal/service"
	"github.com/test/pkg/api"
	"github.com/test/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer()
	api.RegisterOrderServiceServer(server, srv)

	if err := server.Serve(lis); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "dsa", zap.Error(err))
	}

}
