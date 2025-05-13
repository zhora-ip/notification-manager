package server

import (
	"context"
	"fmt"
	"os"

	token "github.com/zhora-ip/notification-manager/pkg/app/verification_token"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/gomail.v2"
)

func (s *Server) VerifyEmail(ctx context.Context, req *ntfs.VerifyEmailRequest) (*ntfs.VerifyEmailResponse, error) {

	eToken := token.GenerateToken(req.GetEmail())
	s.tokens[eToken.Token] = eToken

	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@example.com")
	m.SetHeader("To", req.GetEmail())
	m.SetHeader("Subject", "Подтверждение email")

	link := fmt.Sprintf("https://localhost:8001/verify?token=%s", eToken.Token)
	body := fmt.Sprintf(`Перейдите по ссылке для подтверждения: <a href="%s">%s</a>`, link, link)

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", port, s.email, os.Getenv("MAIL_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return nil, status.Error(codes.Internal, "Failed to send email")
	}

	return &ntfs.VerifyEmailResponse{Success: true, Message: "Verification email sent"}, nil
}
