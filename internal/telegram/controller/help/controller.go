package user

import (
	"context"
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
	c.sender.SendMessage(msg.Chat.ID, &message.Help{
		Message: `Usage

	Domains: /me /sub /help

	/me commands:
		/me subs
			shows list of user subscriptions and its statuses
			no args

	/sub commands:
		/sub add
			adds new subscription, sets unpaid
			args: 
				name - name of service(string with no spaces)
				cost - monthly paid price(positive int)
				cap -  capacity (positive int)
				payday - positive int not grater then 31
			usage pattern: /sub add name=my_awesome_sub cost=500 cap=5 payday=15

		/sub status
			shows status of the subscription, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub status 5

		/sub edit
			edits the subscription, executed only by owner
			args:
				id - id of subscription(implicit)
				name - name of service(string with no spaces)(optional)
				cost - monthly paid price(positive int)(optional)
				cap -  capacity (positive int)(optional)
				payday - positive int not grater then 31(optional)
			usage pattern: /sub edit 5 name=new_awesome_name

		/sub drop
			drops the subscription, executed only by owner
			args:
				id - id of subscription(implicit)
			usage pattern: /sub drop 5

		/sub join
			enter to subscription membership
			args:
				id - id of subscription(implicit)
			usage pattern: /sub join 5

		/sub leave
			leave subscription membership, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub leave 5

		/sub pay
			marks that member is paid his share, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub pay 5

		/sub reset
			resets payment statuses of members, executed only by owner
			args:
				id - id of subscription(implicit)
			usage pattern: /sub reset 5

	/help commands:
		just call this domain to get manual
	`}, true)
}
