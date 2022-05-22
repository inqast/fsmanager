package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) ReadSubscription(ctx context.Context, Id int) (*models.Subscription, error) {

	msg := api.ID{
		Id: int64(Id),
	}

	resp, err := s.grpcClient.ReadSubscription(ctx, &msg)
	if err != nil {
		return &models.Subscription{}, err
	}

	CreatedAt, err := time.Parse(time.RFC3339, resp.CreatedAt)
	if err != nil {
		return &models.Subscription{}, err
	}

	return &models.Subscription{
		ID:                int(resp.Id),
		ChatID:            int(resp.ChatId),
		ServiceName:       resp.ServiceName,
		Capacity:          int(resp.Capacity),
		PriceInCentiUnits: int(resp.PriceInCentiUnits),
		PaymentDay:        int(resp.PaymentDay),
		CreatedAt:         sql.NullTime{Time: CreatedAt},
	}, err
}
