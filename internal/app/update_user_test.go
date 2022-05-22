package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testName := "testName"
	testTelegramId := 12
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateUserMock.Return(nil)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateUser(ctx, &api.User{
		Id:         int64(testId),
		TelegramId: int64(testTelegramId),
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime.Time.Format(time.RFC3339),
	})

	assert.Nil(t, err)
}

func TestUpdateUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateUserMock.Return(repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateUser(ctx, &api.User{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
