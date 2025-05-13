package server

import (
	"context"
	"time"

	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

func (s *Server) ConfirmEmail(ctx context.Context, req *ntfs.ConfirmationRequest) (*ntfs.ConfirmationResponse, error) {
	token, exists := s.tokens[req.Token]
	if !exists {
		return &ntfs.ConfirmationResponse{Verified: false, Message: "Invalid token"}, nil
	}

	if time.Now().After(token.ExpiresAt) {
		return &ntfs.ConfirmationResponse{Verified: false, Message: "Token expired"}, nil
	}

	delete(s.tokens, req.Token)
	return &ntfs.ConfirmationResponse{Verified: true, Message: "Email verified"}, nil
}
