package main

import (
	"context"
	"log"
	"time"

	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	address     = "localhost:50001"
	inputString = "vv_gridin@mail.ru"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := ntfs.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = c.Notify(ctx, &ntfs.NotifyRequest{Email: inputString})
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			log.Fatalf("err")
		}
		log.Fatalf("err")
	}
}
