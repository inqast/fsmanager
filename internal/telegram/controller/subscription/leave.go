package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) leave(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Success, error) {
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
	if subscription.Capacity <= 1 {
		return nil, ErrLastMember
	}

	subscriberId := 0
	isOwner := false
	for _, subscriber := range subscribers {
		if subscriber.UserID == user.ID {
			subscriberId = subscriber.ID
		}
		if subscriber.IsOwner && subscriber.UserID == user.ID {
			isOwner = true
		}
	}
	if subscriberId == 0 {
		return nil, ErrNotListed
	}
	if isOwner {
		return nil, ErrIsOwner
	}

	_, err = c.service.DeleteSubscriber(ctx, int64(subscriberId))
	if err != nil {
		return nil, ErrDeleteFail
	}

	return &message.Success{}, nil
}
