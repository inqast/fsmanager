package authenticator

import (
	"context"
	"errors"
	"log"

	"github.com/inqast/fsmanager/internal/models"
)

func (a *Authenticator) CreateUser(ctx context.Context, telegramId int, name string) (*models.User, error) {
	id, err := a.service.CreateUser(ctx, &models.User{
		Name:       name,
		TelegramID: telegramId,
	})
	if err != nil {
		log.Printf("%+v\n", err)
		return &models.User{}, errors.New("authentication failed")
	}
	return &models.User{
		ID:         id,
		Name:       name,
		TelegramID: telegramId,
	}, err
}
