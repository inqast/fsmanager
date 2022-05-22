package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) pay(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Success, error) {
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

	subscriber := &models.Subscriber{}
	for _, surSubscriber := range subscribers {
		if surSubscriber.UserID == user.ID {
			subscriber = surSubscriber
		}
	}

	if subscriber.ID == 0 {
		return nil, ErrNotListed
	} else if subscriber.IsPaid == true {
		return nil, ErrIsPaid
	}

	subscriber.IsPaid = true
	_, err = c.service.UpdateSubscriber(ctx, subscriber)
	if err != nil {
		return nil, ErrUpdateFail
	}

	return &message.Success{}, nil
}
