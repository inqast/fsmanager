package repository

import (
	"context"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) UpdateUser(ctx context.Context, user models.User) (err error) {

	const query = `
		update users
		set	name = $2,
			pwd = $3,
			telegram_id = $4
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		user.ID,
		user.Name,
		user.Pwd,
		user.TelegramID,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
