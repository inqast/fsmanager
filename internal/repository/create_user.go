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
			created_at
		) VALUES (
			$1, $2, now()
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		user.Name,
		user.Pwd,
	).Scan(&ID)

	return
}
