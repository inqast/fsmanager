package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersByIDs(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testData := []*api.User{
		{
			Id:         1,
			TelegramId: 111,
			Name:       "testName",
			Pwd:        "testPass",
			CreatedAt:  time.Now().Format(time.RFC3339),
		},
		{
			Id:         2,
			TelegramId: 222,
			Name:       "testName2",
			Pwd:        "testPass@",
			CreatedAt:  time.Now().Format(time.RFC3339),
		},
	}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetUsersByIDsMock.Return(&api.GetUsersByIDsResponse{Users: testData}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	users, err := svc.GetUsersByIDs(ctx, []int{1, 2})

	assert.Nil(t, err)
	for i, user := range users {
		testUser := testData[i]
		assert.Equal(t, user.ID, int(testUser.Id))
		assert.Equal(t, user.TelegramID, int(testUser.TelegramId))
		assert.Equal(t, user.Name, testUser.Name)
		assert.Equal(t, user.Pwd, "")
		assert.Equal(t, user.CreatedAt.Time.Format(time.RFC3339), user.CreatedAt.Time.Format(time.RFC3339))
	}
}

func TestGetUsersByIDsNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetUsersByIDsMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.GetUsersByIDs(ctx, []int{1})

	assert.Equal(t, err, ErrNotFound)
}
