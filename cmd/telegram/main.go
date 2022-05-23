package main

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/authenticator"

	grpcClient "github.com/inqast/fsmanager/internal/telegram/grpc"
	"github.com/inqast/fsmanager/internal/telegram/sender"

	helpController "github.com/inqast/fsmanager/internal/telegram/controller/help"
	subscriptionController "github.com/inqast/fsmanager/internal/telegram/controller/subscription"
	userController "github.com/inqast/fsmanager/internal/telegram/controller/user"

	"github.com/inqast/fsmanager/internal/config"

	"github.com/inqast/fsmanager/internal/telegram/router"
	"github.com/inqast/fsmanager/pkg/api"

	"google.golang.org/grpc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.GetConfigFromFile()

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(cfg.Telegram.Offset)
	u.Timeout = cfg.Telegram.Timeout

	updates := bot.GetUpdatesChan(u)
	conn, err := grpc.Dial(cfg.Grpc.Address(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewFamilySubClient(conn)
	grpcService := grpcClient.New(client)

	senderService := sender.New(bot)

	auth := authenticator.New(grpcService)

	uControl := userController.New(senderService, grpcService, auth)
	sControl := subscriptionController.New(senderService, grpcService, auth)
	hControl := helpController.New(senderService, grpcService, auth)

	routerService := router.New(uControl, sControl, hControl)
	ctx := context.Background()
	for update := range updates {
		routerService.HandleUpdate(ctx, update)
	}
}
