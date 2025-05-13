package server

import (
	token "github.com/zhora-ip/notification-manager/pkg/app/verification_token"
	notifications "github.com/zhora-ip/notification-manager/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	email = "libmanager2025@gmail.com"
)

type Server struct {
	tokens map[string]*token.VerificationToken
	email  string
	logger *zap.SugaredLogger
	notifications.UnimplementedNotificationServiceServer
}

func Register(gRPC *grpc.Server, logger *zap.SugaredLogger) {
	notifications.RegisterNotificationServiceServer(gRPC, &Server{
		tokens: map[string]*token.VerificationToken{},
		email:  email,
		logger: logger,
	})
}
