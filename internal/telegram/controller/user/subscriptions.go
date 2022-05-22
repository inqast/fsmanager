package user

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/telegram/message"
)

func (c *Controller) subscriptions(ctx context.Context, chatId, userTelegramId int, userName string) (*message.SubscriptionsResponse, error) {
	user, err := c.auth.Authenticate(ctx, userTelegramId, userName)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	subscriptions, err := c.service.GetSubscriptionsForUser(ctx, user.ID)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, ErrReadFail
	}

	resp := &message.SubscriptionsResponse{
		UserName:      user.Name,
		Subscriptions: make([]*message.Subscription, 0),
	}

	for _, s := range subscriptions {
		if s.ChatID != chatId {
			continue
		}
		subscribers, err := c.service.GetSubscribers(ctx, s.ID)
		if err != nil || len(subscribers) == 0 {
			log.Printf("%+v\n", err)
			return nil, ErrReadFail
		}

		var owner *message.Member
		members := make([]*message.Member, 0)
		for _, subscriber := range subscribers {
			curUser, err := c.service.ReadUser(ctx, subscriber.UserID)
			if err != nil {
				log.Printf("%+v\n", err)
				return nil, ErrReadFail
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

		resp.Subscriptions = append(resp.Subscriptions, &message.Subscription{
			Id:         s.ID,
			Service:    s.ServiceName,
			Owner:      owner.Name,
			Cost:       float64(s.PriceInCentiUnits) / 100,
			PaymentDay: s.PaymentDay,
			Members:    members,
			Capacity:   s.Capacity,
			Share:      float64(s.PriceInCentiUnits/len(members)) / 100,
			IsPaid:     owner.IsPaid,
		})
	}
	return resp, nil
}
