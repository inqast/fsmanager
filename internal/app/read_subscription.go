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

func (t *tserver) ReadSubscription(ctx context.Context, req *api.ID) (*api.Subscription, error) {

	subscription, err := t.repo.ReadSubscription(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.Subscription{
		Id:                int64(subscription.ID),
		OwnerID:           int64(subscription.OwnerID),
		ServiceName:       subscription.ServiceName,
		Capacity:          int64(subscription.Capacity),
		PriceInCentiUnits: int64(subscription.PriceInCentiUnits),
		PaymentDate:       subscription.PaymentDate.Time.Format(time.RFC3339),
		CreatedAt:         subscription.CreatedAt.Time.Format(time.RFC3339),
	}, err
}
