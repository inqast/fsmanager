package repository

import (
	"context"
	"errors"

	"github.com/inqast/fsmanager/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadUser(ctx context.Context, ID int) (user models.User, err error) {
	const query = `
		select id,
			   name,
			   pwd,
			   created_at
		  from users
		where id = $1;
	`
	err = r.pool.QueryRow(ctx, query, ID).Scan(
		&user.ID,
		&user.Name,
		&user.Pwd,
		&user.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}

	return
}
