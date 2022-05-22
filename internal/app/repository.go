package app

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

type Repository interface {
	CreateUser(context.Context, models.User) (int, error)
	ReadUser(context.Context, int) (models.User, error)
	UpdateUser(context.Context, models.User) error
	DeleteUser(context.Context, int) error
	GetUsersByIDs(context.Context, []int64) ([]models.User, error)
	GetUserByTelegramID(context.Context, int) (models.User, error)

	CreateSubscription(context.Context, models.Subscription) (int, error)
	ReadSubscription(context.Context, int) (models.Subscription, error)
	UpdateSubscription(context.Context, models.Subscription) error
	DeleteSubscription(context.Context, int) error

	CreateSubscriber(context.Context, models.Subscriber) (int, error)
	ReadSubscriber(context.Context, int) (models.Subscriber, error)
	UpdateSubscriber(context.Context, models.Subscriber) error
	DeleteSubscriber(context.Context, int) error

	GetSubscribersForSubscription(context.Context, int) ([]models.Subscriber, error)
	GetSubscriptionsForUser(context.Context, int) ([]models.Subscription, error)
}
