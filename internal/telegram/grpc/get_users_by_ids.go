package grpc

import (
	"context"
	"database/sql"
	"time"

	"github.com/inqast/fsmanager/pkg/api"

	"github.com/inqast/fsmanager/internal/models"
)

func (s *Service) GetUsersByIDs(ctx context.Context, ids []int) ([]*models.User, error) {

	idsForRequest := make([]int64, len(ids))
	for i, id := range ids {
		idsForRequest[i] = int64(id)
	}

	msg := api.GetUsersByIDsRequest{
		Ids: idsForRequest,
	}

	resp, err := s.grpcClient.GetUsersByIDs(ctx, &msg)
	if err != nil {
		return []*models.User{}, err
	}

	users := make([]*models.User, len(resp.Users))
	for i, user := range resp.Users {
		createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
		if err != nil {
			return []*models.User{}, err
		}

		users[i] = &models.User{
			ID:         int(user.Id),
			Name:       user.Name,
			TelegramID: int(user.TelegramId),
			CreatedAt:  sql.NullTime{Time: createdAt},
		}
	}

	return users, err
}
