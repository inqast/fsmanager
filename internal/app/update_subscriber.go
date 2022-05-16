package app

import (
	"context"
	"errors"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) UpdateSubscriber(ctx context.Context, req *api.Subscriber) (*emptypb.Empty, error) {

	var subscriber = models.Subscriber{
		ID:             int(req.Id),
		UserID:         int(req.UserID),
		SubscriptionID: int(req.SubscriptionID),
		IsPaid:         req.IsPaid,
		IsOwner:        req.IsOwner,
	}

	err := t.repo.UpdateSubscriber(ctx, subscriber)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, err
}
