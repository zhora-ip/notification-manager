package server

import (
	notifications "github.com/zhora-ip/notification-manager/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	logger *zap.SugaredLogger
	notifications.UnimplementedNotificationServiceServer
}

func Register(gRPC *grpc.Server, logger *zap.SugaredLogger) {
	notifications.RegisterNotificationServiceServer(gRPC, &Server{
		logger: logger,
	})
}
