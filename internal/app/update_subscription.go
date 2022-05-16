package app

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) UpdateSubscription(ctx context.Context, req *api.Subscription) (*emptypb.Empty, error) {
	paymentDate, err := time.Parse(time.RFC3339, req.PaymentDate)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	var subscription = models.Subscription{
		ID:                int(req.Id),
		OwnerID:           int(req.OwnerID),
		ServiceName:       req.ServiceName,
		Capacity:          int(req.Capacity),
		PriceInCentiUnits: int(req.PriceInCentiUnits),
		PaymentDate:       sql.NullTime{Time: paymentDate},
	}

	err = t.repo.UpdateSubscription(ctx, subscription)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, err
}
