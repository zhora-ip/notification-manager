package server

import (
	"context"
	"fmt"
	"os"

	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/gomail.v2"
)

const (
	port = 587
)

func (s *Server) Notify(ctx context.Context, req *ntfs.NotifyRequest) (*ntfs.NotifyResponse, error) {

	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@example.com")
	m.SetHeader("To", req.GetEmail())
	m.SetHeader("Subject", fmt.Sprintf("Заказ №%d", req.GetOrderId()))

	body := fmt.Sprintf(`Добрый день, %s! Ваш заказ №%d просрочен, пожалуйста, верните книгу в библиотеку!`, req.GetName(), req.GetOrderId())
	if req.GetType() == ntfs.NotificationType_ACCEPTED {
		body = fmt.Sprintf(`Добрый день, %s! Заказ №%d ожидает подтверждения!`, req.GetName(), req.GetOrderId())
	}
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", port, s.email, os.Getenv("MAIL_PASSWORD"))

	s.logger.Infow("input", "email", req.GetEmail())

	if err := d.DialAndSend(m); err != nil {
		s.logger.Errorw("send error", "err", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return nil, nil
}
