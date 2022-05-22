package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) UpdateSubscriber(ctx context.Context, subscriber *models.Subscriber) (bool, error) {

	msg := api.Subscriber{
		Id:             int64(subscriber.ID),
		UserID:         int64(subscriber.UserID),
		SubscriptionID: int64(subscriber.SubscriptionID),
		IsPaid:         subscriber.IsPaid,
		IsOwner:        subscriber.IsOwner,
	}

	_, err := s.grpcClient.UpdateSubscriber(ctx, &msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
