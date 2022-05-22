package repository

import (
	"context"
	"errors"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadSubscription(ctx context.Context, ID int) (subscription models.Subscription, err error) {
	const query = `
		select id,
			chat_id,
			service_name,
			capacity,
			price_in_centi_units,
			payment_day,
			created_at
		  from subscriptions
		where id = $1;
	`
	err = r.pool.QueryRow(ctx, query, ID).Scan(
		&subscription.ID,
		&subscription.ChatID,
		&subscription.ServiceName,
		&subscription.Capacity,
		&subscription.PriceInCentiUnits,
		&subscription.PaymentDay,
		&subscription.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}

	return
}
