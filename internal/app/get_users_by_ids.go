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

func (t *tserver) GetUsersByIDs(ctx context.Context, req *api.GetUsersByIDsRequest) (*api.GetUsersByIDsResponse, error) {

	users, err := t.repo.GetUsersByIDs(ctx, req.Ids)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalUsers := make([]*api.User, 0)
	for _, user := range users {
		externalUsers = append(externalUsers, &api.User{
			Id:         int64(user.ID),
			Name:       user.Name,
			TelegramId: int64(user.TelegramID),
			CreatedAt:  user.CreatedAt.Time.Format(time.RFC3339),
		})
	}

	return &api.GetUsersByIDsResponse{
		Users: externalUsers,
	}, err
}
