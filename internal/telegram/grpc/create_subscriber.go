package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) CreateSubscriber(ctx context.Context, subscriber *models.Subscriber) (int, error) {

	msg := api.Subscriber{
		UserID:         int64(subscriber.UserID),
		SubscriptionID: int64(subscriber.SubscriptionID),
		IsPaid:         subscriber.IsPaid,
		IsOwner:        subscriber.IsOwner,
	}

	resp, err := s.grpcClient.CreateSubscriber(ctx, &msg)
	if err != nil {
		return 0, err
	}

	return int(resp.GetId()), nil
}
