package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) GetSubscriptions(ctx context.Context, userID int) (subscriptions []models.Subscription, err error) {
	const query = `
		select subscriptions.id,
			subscriptions.owner_id,
			subscriptions.service_name,
			subscriptions.capacity,
			subscriptions.price_in_centi_units,
			subscriptions.payment_date,
			subscriptions.created_at
	  	from subscribers
		left join subscriptions on  subscribers.subscription_id = subscriptions.id
		where subscribers.user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var subscription models.Subscription
		if err = rows.Scan(
			&subscription.ID,
			&subscription.OwnerID,
			&subscription.ServiceName,
			&subscription.Capacity,
			&subscription.PriceInCentiUnits,
			&subscription.PaymentDate,
			&subscription.CreatedAt,
		); err != nil {
			return
		}

		subscriptions = append(subscriptions, subscription)
	}

	return
}
