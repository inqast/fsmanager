package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) CreateUser(ctx context.Context, User *models.User) (int, error) {

	msg := api.User{
		Name:       User.Name,
		TelegramId: int64(User.TelegramID),
	}

	resp, err := s.grpcClient.CreateUser(ctx, &msg)
	if err != nil {
		return 0, err
	}

	return int(resp.GetId()), nil
}
