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

func (t *tserver) ReadUser(ctx context.Context, req *api.ID) (*api.User, error) {

	user, err := t.repo.ReadUser(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.User{
		Id:         int64(user.ID),
		Name:       user.Name,
		Pwd:        user.Pwd,
		TelegramId: int64(user.TelegramID),
		CreatedAt:  user.CreatedAt.Time.Format(time.RFC3339),
	}, err
}
