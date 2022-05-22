package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/inqast/fsmanager/internal/telegram/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/inqast/fsmanager/internal/telegram/command"
)

type Sender interface {
	SendMessage(int64, fmt.Stringer, bool)
}

type Auth interface {
	Authenticate(context.Context, int, string) (*models.User, error)
}

type GrpcClient interface {
	CreateUser(context.Context, *models.User) (int, error)
	ReadUser(context.Context, int) (*models.User, error)
	UpdateUser(context.Context, *models.User) (bool, error)
	DeleteUser(context.Context, int64) (bool, error)

	CreateSubscription(context.Context, *models.Subscription) (int, error)
	ReadSubscription(context.Context, int) (*models.Subscription, error)
	UpdateSubscription(context.Context, *models.Subscription) (bool, error)
	DeleteSubscription(context.Context, int64) (bool, error)

	CreateSubscriber(context.Context, *models.Subscriber) (int, error)
	ReadSubscriber(context.Context, int) (*models.Subscriber, error)
	UpdateSubscriber(context.Context, *models.Subscriber) (bool, error)
	DeleteSubscriber(context.Context, int64) (bool, error)

	GetByTelegramID(context.Context, int) (*models.User, error)
	GetSubscribers(context.Context, int) ([]*models.Subscriber, error)
	GetSubscriptionsForUser(context.Context, int) ([]*models.Subscription, error)
	GetUsersByIDs(context.Context, []int) ([]*models.User, error)
}

var ErrReadFail = errors.New("read failed")

type Controller struct {
	sender  Sender
	service GrpcClient
	auth    Auth
}

func New(
	sender Sender,
	service GrpcClient,
	auth Auth,
) *Controller {
	return &Controller{
		sender:  sender,
		service: service,
		auth:    auth,
	}
}

func (c *Controller) HandleCommand(ctx context.Context, msg *tgbotapi.Message, query *command.Query) {
	var response fmt.Stringer
	var err error

	switch query.Command {
	case "subs":
		response, err = c.subscriptions(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
	default:
		response = &message.Error{Msg: "command not found"}
	}

	c.sender.SendMessage(msg.Chat.ID, response, true)
}
