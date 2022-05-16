package repository

import (
	"context"
	"time"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) UpdateSubscription(ctx context.Context, subscription models.Subscription) (err error) {

	const query = `
		update subscriptions
		set	owner_id = $2,
			service_name = $3,
			capacity = $4,
			price_in_centi_units = $5,
			payment_date = $6
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		subscription.ID,
		subscription.OwnerID,
		subscription.ServiceName,
		subscription.Capacity,
		subscription.PriceInCentiUnits,
		subscription.PaymentDate.Time.Format(time.RFC3339),
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
