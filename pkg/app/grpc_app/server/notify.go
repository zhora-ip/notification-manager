package server

import (
	"context"
	"os"

	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/gomail.v2"
)

func (s *Server) Notify(ctx context.Context, req *ntfs.NotifyRequest) (*ntfs.NotifyResponse, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "libmanager2025@gmail.com")
	m.SetHeader("To", req.GetEmail())
	m.SetHeader("Subject", "Тест")
	m.SetBody("text/html", "<b>Привет</b> <i>мир!</i>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "libmanager2025@gmail.com", os.Getenv("MAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return nil, nil
}
