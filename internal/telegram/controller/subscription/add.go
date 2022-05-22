package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"

	"github.com/inqast/fsmanager/internal/models"
)

func (c *Controller) add(ctx context.Context, chatId, telegramId int, name string, args []string) (*message.Success, error) {
	user, err := c.auth.Authenticate(ctx, telegramId, name)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	subscription, err := c.parseArgs(args)
	isValid := c.validateModel(subscription)
	if err != nil || !isValid {
		log.Print(err)
		return nil, ErrIncorrectArgs
	}
	subscription.ChatID = chatId

	id, err := c.service.CreateSubscription(ctx, subscription)
	if err != nil {
		log.Print(err)
		return nil, ErrCreationFail
	}

	_, err = c.service.CreateSubscriber(ctx, &models.Subscriber{
		UserID:         user.ID,
		SubscriptionID: id,
		IsOwner:        true,
		IsPaid:         false,
	})
	if err != nil {
		log.Print(err)
		return nil, ErrCreationFail
	}

	return &message.Success{Id: id}, nil
}

func (c Controller) validateModel(subscription *models.Subscription) bool {
	switch {
	case subscription.ServiceName == "":
		return false
	case subscription.PriceInCentiUnits == 0:
		return false
	case subscription.Capacity == 0:
		return false
	case subscription.PaymentDay == 0:
		return false
	}

	return true
}
