package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) CreateSubscriber(ctx context.Context, subscriber models.Subscriber) (ID int, err error) {

	const query = `
		insert into subscribers (
			user_id,
			subscription_id,
			is_paid,
			is_owner,
			created_at
		) VALUES (
			$1, $2, $3, $4, now()
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		subscriber.UserID,
		subscriber.SubscriptionID,
		subscriber.IsPaid,
		subscriber.IsOwner,
	).Scan(&ID)

	return
}
