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

func TestGetSubscribers(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testSubscriptionId := 2

	testData := []models.Subscriber{
		{
			ID:             1,
			UserID:         2,
			SubscriptionID: 3,
			IsPaid:         false,
			IsOwner:        true,
			CreatedAt:      sql.NullTime{Time: time.Now()},
		},
		{
			ID:             2,
			UserID:         1,
			SubscriptionID: 3,
			IsPaid:         false,
			IsOwner:        false,
			CreatedAt:      sql.NullTime{Time: time.Now()},
		},
	}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetSubscribersMock.Return(testData, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	subscribers, err := svc.GetSubscribers(ctx, &api.ID{Id: int64(testSubscriptionId)})

	assert.Nil(t, err)
	for i, subscriber := range subscribers.Subscribers {
		testSubscriber := testData[i]
		assert.Equal(t, subscriber.Id, int64(testSubscriber.ID))
		assert.Equal(t, subscriber.UserID, int64(testSubscriber.UserID))
		assert.Equal(t, subscriber.SubscriptionID, int64(testSubscriber.SubscriptionID))
		assert.Equal(t, subscriber.IsPaid, testSubscriber.IsPaid)
		assert.Equal(t, subscriber.IsOwner, testSubscriber.IsOwner)
		assert.Equal(t, subscriber.CreatedAt, testSubscriber.CreatedAt.Time.Format(time.RFC3339))
	}
}

func TestGetSubscribersNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetSubscribersMock.Return([]models.Subscriber{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.GetSubscribers(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
