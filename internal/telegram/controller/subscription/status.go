package subscription

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) status(ctx context.Context, chatId, userTelegramId int, userName string, args []string) (*message.Subscription, error) {
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

	var subscriberId int
	var owner *message.Member
	members := make([]*message.Member, 0)
	for _, subscriber := range subscribers {
		curUser, err := c.service.ReadUser(ctx, subscriber.UserID)
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, ErrReadFail
		}
		if curUser.ID == user.ID {
			subscriberId = subscriber.ID
		}
		member := message.Member{
			UserID:  subscriber.UserID,
			Name:    curUser.Name,
			IsOwner: subscriber.IsOwner,
			IsPaid:  subscriber.IsPaid,
		}

		members = append(members, &member)
		if subscriber.IsOwner {
			owner = &member
		}
	}
	if subscriberId == 0 {
		return nil, ErrNotListed
	}

	return &message.Subscription{
		Id:         subscription.ID,
		Service:    subscription.ServiceName,
		Owner:      owner.Name,
		Cost:       float64(subscription.PriceInCentiUnits) / 100,
		PaymentDay: subscription.PaymentDay,
		Members:    members,
		Capacity:   subscription.Capacity,
		Share:      float64(subscription.PriceInCentiUnits/len(members)) / 100,
		IsPaid:     owner.IsPaid,
	}, nil
}
