package app

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
)

func (t *tserver) CreateUser(ctx context.Context, req *api.User) (*api.ID, error) {

	var user = models.User{
		Name:       req.Name,
		Pwd:        req.Pwd,
		TelegramID: int(req.TelegramId),
	}

	ID, err := t.repo.CreateUser(ctx, user)

	return &api.ID{
		Id: int64(ID),
	}, err
}
