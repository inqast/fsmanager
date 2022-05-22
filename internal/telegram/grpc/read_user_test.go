package grpc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestReadUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testChatID := 2
	testName := "testName"
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadUserMock.Return(&api.User{
		Id:         int64(testId),
		TelegramId: int64(testChatID),
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime.Time.Format(time.RFC3339),
	}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	user, err := svc.ReadUser(ctx, testId)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, testId)
	assert.Equal(t, user.TelegramID, testChatID)
	assert.Equal(t, user.Name, testName)
	assert.Equal(t, user.Pwd, "")
	assert.Equal(t, user.CreatedAt.Time.Format(time.RFC3339), testTime.Time.Format(time.RFC3339))
}

func TestReadUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadUserMock.Return(&api.User{}, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.ReadUser(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
