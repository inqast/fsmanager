package subscription

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

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

var ErrIncorrectArgs = errors.New("incorrect args, try name={string} cost={int} cap={int} payday={int}")
var ErrIncorrectArgsCount = errors.New("wrong number of arguments")
var ErrCreationFail = errors.New("creation failed")
var ErrOutOfGroup = errors.New("subscription if out of group")
var ErrLastMember = errors.New("you are last member")
var ErrNotListed = errors.New("you are not listed")
var ErrNotOwner = errors.New("you are not owner")
var ErrDeleteFail = errors.New("delete failed")
var ErrCapacityLess = errors.New("capacity can't be less then members count")
var ErrUpdateFail = errors.New("update failed")
var ErrCapacityOverload = errors.New("capacity overload")
var ErrListed = errors.New("you are already listed")
var ErrIsOwner = errors.New("you are owner")
var ErrReadFail = errors.New("read failed")
var ErrIsPaid = errors.New("already paid")

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
	case "add":
		response, err = c.add(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "join":
		response, err = c.join(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "leave":
		response, err = c.leave(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "drop":
		response, err = c.drop(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "pay":
		response, err = c.pay(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "edit":
		response, err = c.edit(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "status":
		response, err = c.status(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	case "notify":
		response, err = c.notify(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, false)
	case "reset":
		response, err = c.reset(ctx, int(msg.Chat.ID), int(msg.From.ID), msg.From.UserName, query.Args)
		if err != nil {
			response = &message.Error{Msg: err.Error()}
		}
		c.sender.SendMessage(msg.Chat.ID, response, true)
	default:
		c.sender.SendMessage(msg.Chat.ID, &message.Error{Msg: "command not found"}, true)
	}
}

func (c *Controller) getData(
	ctx context.Context,
	userTelegramId int,
	userName,
	subscriptionIdArg string) (*models.User, *models.Subscription, []*models.Subscriber, error) {
	user, err := c.auth.Authenticate(ctx, userTelegramId, userName)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, nil, nil, err
	}

	subscriptionId, err := strconv.Atoi(subscriptionIdArg)
	if err != nil {
		return nil, nil, nil, errors.New("wrong type of argument use int")
	}
	subscription, err := c.service.ReadSubscription(ctx, subscriptionId)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, nil, nil, errors.New("subscription read failed")
	}

	subscribers, err := c.service.GetSubscribers(ctx, subscription.ID)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, nil, nil, errors.New("subscribers read failed")
	}

	return user, subscription, subscribers, nil
}
