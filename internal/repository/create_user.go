package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) CreateUser(ctx context.Context, user models.User) (ID int, err error) {

	const query = `
		insert into users (
			name,
			pwd,
			telegram_id,
			created_at
		) VALUES (
			$1, $2, $3, now()
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		user.Name,
		user.Pwd,
		user.TelegramID,
	).Scan(&ID)

	return
}
