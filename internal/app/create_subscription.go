package app

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (t *tserver) CreateSubscription(ctx context.Context, req *api.Subscription) (*api.ID, error) {

	paymentDate, err := time.Parse(time.RFC3339, req.PaymentDate)
	if err != nil {
		return &api.ID{
			Id: int64(0),
		}, err
	}

	var subscription = models.Subscription{
		OwnerID:           int(req.OwnerID),
		ServiceName:       req.ServiceName,
		Capacity:          int(req.Capacity),
		PriceInCentiUnits: int(req.PriceInCentiUnits),
		PaymentDate:       sql.NullTime{Time: paymentDate},
	}

	ID, err := t.repo.CreateSubscription(ctx, subscription)

	return &api.ID{
		Id: int64(ID),
	}, err
}
