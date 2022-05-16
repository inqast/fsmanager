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

func TestReadUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testName := "testName"
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadUserMock.Return(models.User{
		ID:        testId,
		Name:      testName,
		Pwd:       testPwd,
		CreatedAt: testTime,
	}, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	user, err := svc.ReadUser(ctx, &api.ID{Id: int64(testId)})

	assert.Nil(t, err)
	assert.Equal(t, user.Id, int64(testId))
	assert.Equal(t, user.Name, testName)
	assert.Equal(t, user.Pwd, testPwd)
	assert.Equal(t, user.CreatedAt, testTime.Time.Format(time.RFC3339))
}

func TestReadUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadUserMock.Return(models.User{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.ReadUser(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
