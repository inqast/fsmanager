package repository

import (
	"context"
	"errors"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadSubscriber(ctx context.Context, ID int) (user models.Subscriber, err error) {
	const query = `
		select id,
			user_id,
			subscription_id,
			is_paid,
			is_owner,
			created_at
		  from subscribers
		where id = $1;
	`
	err = r.pool.QueryRow(ctx, query, ID).Scan(
		&user.ID,
		&user.UserID,
		&user.SubscriptionID,
		&user.IsPaid,
		&user.IsOwner,
		&user.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}

	return
}
