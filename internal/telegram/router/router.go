package router

import (
	"context"
	"log"
	"runtime/debug"

	"github.com/inqast/fsmanager/internal/telegram/command"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller interface {
	HandleCommand(context.Context, *tgbotapi.Message, *command.Query)
}

type Router struct {
	userController         Controller
	subscriptionController Controller
	helpController         Controller
}

func New(userController, subscriptionController, helpController Controller) *Router {
	return &Router{
		userController:         userController,
		subscriptionController: subscriptionController,
		helpController:         helpController,
	}
}

func (r *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	if !update.Message.IsCommand() {
		r.helpController.HandleCommand(ctx, update.Message, &command.Query{})
		return
	}

	query, err := command.Parse(update.Message.Text)
	if err != nil {
		r.helpController.HandleCommand(ctx, update.Message, &command.Query{})
		return
	}

	switch query.Domain {
	case "/me":
		r.userController.HandleCommand(ctx, update.Message, query)
	case "/sub":
		r.subscriptionController.HandleCommand(ctx, update.Message, query)
	default:
		r.helpController.HandleCommand(ctx, update.Message, query)
	}
}
