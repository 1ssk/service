package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/1ssk/service/integnal/config"
	"github.com/1ssk/service/integnal/service"
	"github.com/1ssk/service/pkg/api"
	"github.com/1ssk/service/pkg/logger"
	"github.com/1ssk/service/pkg/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	ctx, _ = logger.New(ctx)

	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to load config", zap.Error(err))
	}

	pool, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect database", zap.Error(err))
	}

	grpsLis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer(grpc.UnaryInterceptor(logger.Interceptor))
	api.RegisterOrderServiceServer(server, srv)

	rt := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterOrderServiceHandlerFromEndpoint(ctx, rt, "localhost:"+strconv.Itoa(cfg.GRPCPort), opts)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to register handler server", zap.Error(err))
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%d", cfg.RestPort),
		Handler: rt,
	}

	httpLis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.RestPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := httpServer.Serve(httpLis); err != nil && err != http.ErrServerClosed {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve HTTP", zap.Error(err))
		}
	}()

	go func() {
		if err := httpServer.Serve(grpsLis); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve", zap.Error(err))
		}
	}()

	select {
	case <-ctx.Done():
		server.GracefulStop()
		pool.Close()
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to shutdown HTTP server", zap.Error(err))
		}
		logger.GetLoggerFromCtx(ctx).Info(ctx, "server stopped")
	}

}
