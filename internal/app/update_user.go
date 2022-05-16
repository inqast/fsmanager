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

func (t *tserver) UpdateUser(ctx context.Context, req *api.User) (*emptypb.Empty, error) {

	var user = models.User{
		ID:   int(req.Id),
		Name: req.Name,
		Pwd:  req.Pwd,
	}

	err := t.repo.UpdateUser(ctx, user)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, err
}
