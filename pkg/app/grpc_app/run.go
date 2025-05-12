package grpcapp

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/zhora-ip/notification-manager/pkg/app/grpc_app/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	serverPort = ":50001"
)

func Run() {
	srv := grpc.NewServer()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	server.Register(srv, sugar)

	errCh := make(chan error, 1)
	go func() {
		log.Print("Server is up and running")
		lis, err := net.Listen("tcp", serverPort)

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		errCh <- srv.Serve(lis)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigCh:
		log.Print("Received terminate, graceful shutdown!")
		srv.Stop()
	case err := <-errCh:
		if err != nil {
			log.Printf("GRPC server error, %v", err)
			return
		}
	}
	log.Print("Server exited gracefully")

}
