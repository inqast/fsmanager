package repository

import "context"

func (r *repository) DeleteSubscriber(ctx context.Context, ID int) (err error) {

	const query = `
		delete from subscribers
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query, ID)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
