package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) UpdateSubscription(ctx context.Context, subscription *models.Subscription) (bool, error) {
	msg := api.Subscription{
		Id:                int64(subscription.ID),
		ChatId:            int64(subscription.ChatID),
		ServiceName:       subscription.ServiceName,
		Capacity:          int64(subscription.Capacity),
		PriceInCentiUnits: int64(subscription.PriceInCentiUnits),
		PaymentDay:        int64(subscription.PaymentDay),
	}

	_, err := s.grpcClient.UpdateSubscription(ctx, &msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
