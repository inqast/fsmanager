package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"

	"github.com/inqast/fsmanager/internal/models"
)

func (r *repository) GetUserByTelegramID(ctx context.Context, telegramID int) (user models.User, err error) {
	const query = `
		select id,
			   name,
			   telegram_id,
			   created_at
	  	from users
		where telegram_id = $1;
	`
	log.Print("log repo")
	err = r.pool.QueryRow(ctx, query, telegramID).Scan(
		&user.ID,
		&user.Name,
		&user.TelegramID,
		&user.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}

	return
}
