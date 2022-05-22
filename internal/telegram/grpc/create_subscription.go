package grpc

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) CreateSubscription(ctx context.Context, subscription *models.Subscription) (int, error) {

	msg := api.Subscription{
		ChatId:            int64(subscription.ChatID),
		ServiceName:       subscription.ServiceName,
		Capacity:          int64(subscription.Capacity),
		PriceInCentiUnits: int64(subscription.PriceInCentiUnits),
		PaymentDay:        int64(subscription.PaymentDay),
	}

	resp, err := s.grpcClient.CreateSubscription(ctx, &msg)
	if err != nil {
		return 0, err
	}

	return int(resp.GetId()), nil
}
