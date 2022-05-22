package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersByIDs(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testData := []models.User{
		{
			ID:         1,
			TelegramID: 111,
			Name:       "testName",
			Pwd:        "testPass",
			CreatedAt:  sql.NullTime{Time: time.Now()},
		},
		{
			ID:         2,
			TelegramID: 222,
			Name:       "testName2",
			Pwd:        "testPass@",
			CreatedAt:  sql.NullTime{Time: time.Now()},
		},
	}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetUsersByIDsMock.Return(testData, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	users, err := svc.GetUsersByIDs(ctx, &api.GetUsersByIDsRequest{
		Ids: []int64{1, 2},
	})

	assert.Nil(t, err)
	for i, user := range users.Users {
		testUser := testData[i]
		assert.Equal(t, user.Id, int64(testUser.ID))
		assert.Equal(t, user.TelegramId, int64(testUser.TelegramID))
		assert.Equal(t, user.Name, testUser.Name)
		assert.Equal(t, user.Pwd, "")
		assert.Equal(t, user.CreatedAt, testUser.CreatedAt.Time.Format(time.RFC3339))
	}
}
