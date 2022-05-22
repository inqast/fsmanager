package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) join(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Success, error) {
	if len(args) != 1 {
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
	if subscription.Capacity <= len(subscribers) {
		return nil, ErrCapacityOverload
	}
	for _, subscriber := range subscribers {
		if subscriber.UserID == user.ID {
			return nil, ErrListed
		}
	}

	_, err = c.service.CreateSubscriber(ctx, &models.Subscriber{
		UserID:         user.ID,
		SubscriptionID: subscription.ID,
		IsPaid:         false,
		IsOwner:        false,
	})
	if err != nil {
		return nil, ErrCreationFail
	}

	return &message.Success{}, nil
}
