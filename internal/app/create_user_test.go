package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testTelegramId := 12131
	testName := "testName"
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateUserMock.Return(testId, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	id, err := svc.CreateUser(ctx, &api.User{
		Id:         int64(testId),
		TelegramId: int64(testTelegramId),
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime.Time.Format(time.RFC3339),
	})

	assert.Nil(t, err)
	assert.Equal(t, id.Id, int64(testId))
}
