package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) GetSubscriptionsForUser(ctx context.Context, id int) ([]*models.Subscription, error) {

	msg := api.ID{
		Id: int64(id),
	}

	resp, err := s.grpcClient.GetSubscriptionsForUser(ctx, &msg)
	if err != nil {
		return []*models.Subscription{}, err
	}

	subscriptions := make([]*models.Subscription, len(resp.Subscriptions))
	for i, subscription := range resp.Subscriptions {
		createdAt, err := time.Parse(time.RFC3339, subscription.CreatedAt)
		if err != nil {
			return []*models.Subscription{}, err
		}

		subscriptions[i] = &models.Subscription{
			ID:                int(subscription.Id),
			ChatID:            int(subscription.ChatId),
			ServiceName:       subscription.ServiceName,
			Capacity:          int(subscription.Capacity),
			PriceInCentiUnits: int(subscription.PriceInCentiUnits),
			PaymentDay:        int(subscription.PaymentDay),
			CreatedAt:         sql.NullTime{Time: createdAt},
		}
	}

	return subscriptions, err
}
