package app

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteUserMock.Return(nil)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.DeleteUser(ctx, &api.ID{Id: int64(testId)})

	assert.Nil(t, err)
}

func TestDeleteUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteUserMock.Return(repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.DeleteUser(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
