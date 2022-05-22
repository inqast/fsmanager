package grpc

import (
	"context"

	"github.com/inqast/fsmanager/pkg/api"
)

func (s *Service) DeleteSubscriber(ctx context.Context, Id int64) (bool, error) {

	msg := api.ID{
		Id: Id,
	}

	_, err := s.grpcClient.DeleteSubscriber(ctx, &msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
