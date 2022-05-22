package repository

import (
	"context"
	"log"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) GetSubscribersForSubscription(ctx context.Context, subscriptionID int) (subscribers []models.Subscriber, err error) {
	const query = `
		select id,
			user_id,
			subscription_id,
			is_paid,
			is_owner,
			created_at
	  	from subscribers
		where subscription_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, subscriptionID)
	if err != nil {
		return
	}
	defer rows.Close()
	log.Println(subscriptionID)
	for rows.Next() {
		var subscriber models.Subscriber
		if err = rows.Scan(
			&subscriber.ID,
			&subscriber.UserID,
			&subscriber.SubscriptionID,
			&subscriber.IsPaid,
			&subscriber.IsOwner,
			&subscriber.CreatedAt,
		); err != nil {
			return
		}

		subscribers = append(subscribers, subscriber)
	}
	log.Println(len(subscribers))

	return
}
