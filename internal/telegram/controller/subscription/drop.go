package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) drop(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Success, error) {
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

	_, err = c.service.DeleteSubscription(ctx, int64(subscription.ID))
	if err != nil {
		return nil, ErrDeleteFail
	}

	return &message.Success{}, nil
}
