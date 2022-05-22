package app

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (t *tserver) CreateSubscription(ctx context.Context, req *api.Subscription) (*api.ID, error) {
	var subscription = models.Subscription{
		ChatID:            int(req.ChatId),
		ServiceName:       req.ServiceName,
		Capacity:          int(req.Capacity),
		PriceInCentiUnits: int(req.PriceInCentiUnits),
		PaymentDay:        int(req.PaymentDay),
	}

	ID, err := t.repo.CreateSubscription(ctx, subscription)

	return &api.ID{
		Id: int64(ID),
	}, err
}
