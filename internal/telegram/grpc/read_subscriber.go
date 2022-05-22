package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) ReadSubscriber(ctx context.Context, Id int) (*models.Subscriber, error) {

	msg := api.ID{
		Id: int64(Id),
	}

	resp, err := s.grpcClient.ReadSubscriber(ctx, &msg)
	if err != nil {
		return &models.Subscriber{}, err
	}

	CreatedAt, err := time.Parse(time.RFC3339, resp.CreatedAt)
	if err != nil {
		return &models.Subscriber{}, err
	}

	return &models.Subscriber{
		ID:             int(resp.Id),
		UserID:         int(resp.UserID),
		SubscriptionID: int(resp.SubscriptionID),
		IsPaid:         resp.IsPaid,
		IsOwner:        resp.IsOwner,
		CreatedAt:      sql.NullTime{Time: CreatedAt},
	}, err
}
