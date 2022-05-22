package grpc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/inqast/fsmanager/internal/models"

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

	clientMock := NewFamilySubClientMock(mc)
	clientMock.CreateUserMock.Return(&api.ID{Id: int64(testId)}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	id, err := svc.CreateUser(ctx, &models.User{
		ID:         testId,
		TelegramID: testTelegramId,
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime,
	})

	assert.Nil(t, err)
	assert.Equal(t, id, testId)
}
