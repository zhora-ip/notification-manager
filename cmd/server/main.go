package main

import (
	"log"

	"github.com/joho/godotenv"
	grpcapp "github.com/zhora-ip/notification-manager/pkg/app/grpc_app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}

func main() {
	grpcapp.Run()
}
