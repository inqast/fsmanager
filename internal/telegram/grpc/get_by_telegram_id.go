package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) GetByTelegramID(ctx context.Context, Id int) (*models.User, error) {
	msg := api.ID{
		Id: int64(Id),
	}

	resp, err := s.grpcClient.GetUserByTelegramID(ctx, &msg)
	if err != nil {
		return &models.User{}, err
	}
	CreatedAt, err := time.Parse(time.RFC3339, resp.CreatedAt)
	if err != nil {
		return &models.User{}, err
	}
	return &models.User{
		ID:         int(resp.Id),
		Name:       resp.Name,
		TelegramID: int(resp.TelegramId),
		CreatedAt:  sql.NullTime{Time: CreatedAt},
	}, err
}
