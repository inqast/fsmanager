package repository

import (
	"context"
	"time"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) CreateSubscription(ctx context.Context, subscription models.Subscription) (ID int, err error) {

	const query = `
		insert into subscriptions (
			owner_id,
			service_name,
			capacity,
			price_in_centi_units,
			payment_date,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, now()
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		subscription.OwnerID,
		subscription.ServiceName,
		subscription.Capacity,
		subscription.PriceInCentiUnits,
		subscription.PaymentDate.Time.Format(time.RFC3339),
	).Scan(&ID)

	return
}
