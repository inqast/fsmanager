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

func TestCreateSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateSubscriberMock.Return(testId, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	subscriber, err := svc.CreateSubscriber(ctx, &api.Subscriber{
		Id:             int64(testId),
		UserID:         int64(testUserID),
		SubscriptionID: int64(testSubscriptionID),
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime.Time.Format(time.RFC3339),
	})

	assert.Nil(t, err)
	assert.Equal(t, subscriber.Id, int64(testId))
}
