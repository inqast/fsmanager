package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) UpdateSubscriber(ctx context.Context, subscriber models.Subscriber) (err error) {

	const query = `
		update subscribers
		set	user_id = $2,
			subscription_id = $3,
			is_paid = $4,
			is_owner = $5
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		subscriber.ID,
		subscriber.UserID,
		subscriber.SubscriptionID,
		subscriber.IsPaid,
		subscriber.IsOwner,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
