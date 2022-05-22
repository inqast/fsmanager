package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) UpdateSubscription(ctx context.Context, subscription models.Subscription) (err error) {

	const query = `
		update subscriptions
		set chat_id = $2,
			service_name = $3,
			capacity = $4,
			price_in_centi_units = $5,
			payment_day = $6
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		subscription.ID,
		subscription.ChatID,
		subscription.ServiceName,
		subscription.Capacity,
		subscription.PriceInCentiUnits,
		subscription.PaymentDay,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
