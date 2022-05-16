package app

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (t *tserver) CreateSubscriber(ctx context.Context, req *api.Subscriber) (*api.ID, error) {

	var subscriber = models.Subscriber{
		UserID:         int(req.UserID),
		SubscriptionID: int(req.SubscriptionID),
		IsPaid:         req.IsPaid,
		IsOwner:        req.IsOwner,
	}

	ID, err := t.repo.CreateSubscriber(ctx, subscriber)

	return &api.ID{
		Id: int64(ID),
	}, err
}
