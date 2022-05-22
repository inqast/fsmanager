package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUserByTelegramID(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testTelegramID := 2
	testName := "testName"
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetUserByTelegramIDMock.Return(models.User{
		ID:         testId,
		TelegramID: testTelegramID,
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime,
	}, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	user, err := svc.GetUserByTelegramID(ctx, &api.ID{Id: int64(testTelegramID)})

	assert.Nil(t, err)
	assert.Equal(t, user.Id, int64(testId))
	assert.Equal(t, user.TelegramId, int64(testTelegramID))
	assert.Equal(t, user.Name, testName)
	assert.Equal(t, user.Pwd, "")
	assert.Equal(t, user.CreatedAt, testTime.Time.Format(time.RFC3339))
}

func TestGetUserByTelegramIDNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetUserByTelegramIDMock.Return(models.User{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.GetUserByTelegramID(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
