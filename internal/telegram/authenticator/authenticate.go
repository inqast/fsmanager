package authenticator

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/inqast/fsmanager/internal/models"
)

func (a *Authenticator) Authenticate(ctx context.Context, telegramId int, name string) (*models.User, error) {
	user, err := a.service.GetByTelegramID(ctx, telegramId)
	if errors.Is(err, status.Error(codes.NotFound, "not found")) {
		return a.CreateUser(ctx, telegramId, name)
	}

	return user, err
}
