package app

import (
	"context"
	"errors"
	"time"

	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) GetSubscriptionsForUser(ctx context.Context, req *api.ID) (*api.GetSubscriptionsResponse, error) {

	subscriptions, err := t.repo.GetSubscriptionsForUser(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalSubscriptions := make([]*api.Subscription, 0)
	for _, subscription := range subscriptions {
		externalSubscriptions = append(externalSubscriptions, &api.Subscription{
			Id:                int64(subscription.ID),
			ChatId:            int64(subscription.ChatID),
			ServiceName:       subscription.ServiceName,
			Capacity:          int64(subscription.Capacity),
			PriceInCentiUnits: int64(subscription.PriceInCentiUnits),
			PaymentDay:        int64(subscription.PaymentDay),
			CreatedAt:         subscription.CreatedAt.Time.Format(time.RFC3339),
		})
	}

	return &api.GetSubscriptionsResponse{
		Subscriptions: externalSubscriptions,
	}, err
}
