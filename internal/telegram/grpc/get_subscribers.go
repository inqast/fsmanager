package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) GetSubscribers(ctx context.Context, id int) ([]*models.Subscriber, error) {

	msg := api.ID{
		Id: int64(id),
	}

	resp, err := s.grpcClient.GetSubscribersForSubscription(ctx, &msg)
	if err != nil {
		return []*models.Subscriber{}, err
	}

	subscribers := make([]*models.Subscriber, len(resp.Subscribers))
	for i, subscriber := range resp.Subscribers {
		createdAt, err := time.Parse(time.RFC3339, subscriber.CreatedAt)
		if err != nil {
			return []*models.Subscriber{}, err
		}

		subscribers[i] = &models.Subscriber{
			ID:             int(subscriber.Id),
			UserID:         int(subscriber.UserID),
			SubscriptionID: int(subscriber.SubscriptionID),
			IsPaid:         subscriber.IsPaid,
			IsOwner:        subscriber.IsOwner,
			CreatedAt:      sql.NullTime{Time: createdAt},
		}
	}

	return subscribers, err
}
