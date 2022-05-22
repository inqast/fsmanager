package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) UpdateUser(ctx context.Context, User *models.User) (bool, error) {

	msg := api.User{
		Id:         int64(User.ID),
		Name:       User.Name,
		TelegramId: int64(User.TelegramID),
	}

	_, err := s.grpcClient.UpdateUser(ctx, &msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
