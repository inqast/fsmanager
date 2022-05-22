package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) edit(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Success, error) {
	if len(args) < 1 {
		return nil, ErrIncorrectArgsCount
	}

	user, subscription, subscribers, err := c.getData(ctx, userTelegramId, userName, args[0])
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	if subscription.ChatID != chatId {
		return nil, ErrOutOfGroup
	}
	isMember := false
	isOwner := false
	for _, subscriber := range subscribers {
		if subscriber.UserID == user.ID {
			isMember = true
			if subscriber.IsOwner {
				isOwner = true
			}
		}
	}
	if !isMember {
		return nil, ErrNotListed
	}
	if !isOwner {
		return nil, ErrNotOwner
	}

	editedSubscription, err := c.parseArgs(args[1:])
	if err != nil {
		log.Print(err)
		return nil, ErrIncorrectArgs
	}

	if editedSubscription.ServiceName != "" {
		subscription.ServiceName = editedSubscription.ServiceName
	}
	if editedSubscription.PriceInCentiUnits != 0 {
		subscription.PriceInCentiUnits = editedSubscription.PriceInCentiUnits
	}
	if editedSubscription.Capacity != 0 {
		if editedSubscription.Capacity < len(subscribers) {
			return nil, ErrCapacityLess
		}
		subscription.Capacity = editedSubscription.Capacity
	}
	if editedSubscription.PaymentDay != 0 {
		subscription.PaymentDay = editedSubscription.PaymentDay
	}

	_, err = c.service.UpdateSubscription(ctx, subscription)
	if err != nil {
		return nil, ErrUpdateFail
	}

	return &message.Success{}, nil
}
