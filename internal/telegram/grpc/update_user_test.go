package grpc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testName := "testName"
	testTelegramId := 12
	testPwd := "testPwd"
	testTime := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateUserMock.Return(nil, nil)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateUser(ctx, &models.User{
		ID:         testId,
		TelegramID: testTelegramId,
		Name:       testName,
		Pwd:        testPwd,
		CreatedAt:  testTime,
	})

	assert.Nil(t, err)
}

func TestUpdateUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateUserMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateUser(ctx, &models.User{ID: 1})

	assert.Equal(t, err, ErrNotFound)
}
