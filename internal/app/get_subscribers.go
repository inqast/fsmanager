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

func (t *tserver) GetSubscribers(ctx context.Context, req *api.ID) (*api.GetSubscribersResponse, error) {

	subscriptions, err := t.repo.GetSubscribers(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalSubscribers := make([]*api.Subscriber, 0)
	for _, subscriber := range subscriptions {
		externalSubscribers = append(externalSubscribers, &api.Subscriber{
			Id:             int64(subscriber.ID),
			UserID:         int64(subscriber.UserID),
			SubscriptionID: int64(subscriber.SubscriptionID),
			IsPaid:         subscriber.IsPaid,
			IsOwner:        subscriber.IsOwner,
			CreatedAt:      subscriber.CreatedAt.Time.Format(time.RFC3339),
		})
	}

	return &api.GetSubscribersResponse{
		Subscribers: externalSubscribers,
	}, err
}
