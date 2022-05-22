package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) GetUsersByIDs(ctx context.Context, IDs []int64) (users []models.User, err error) {
	const query = `
		select id,
			   name,
			   telegram_id,
			   created_at
	  	from users
		where telegram_id = ANY($1);
	`
	rows, err := r.pool.Query(ctx, query, IDs)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.TelegramID,
			&user.CreatedAt,
		); err != nil {
			return
		}

		users = append(users, user)
	}

	return
}
