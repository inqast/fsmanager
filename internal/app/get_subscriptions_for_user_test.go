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

func TestGetSubscriptions(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserID := 2

	testData := []models.Subscription{
		{
			ID:                1,
			ChatID:            2,
			ServiceName:       "testName",
			Capacity:          5,
			PriceInCentiUnits: 500,
			PaymentDay:        12,
			CreatedAt:         sql.NullTime{Time: time.Now()},
		},
		{
			ID:                2,
			ChatID:            1,
			ServiceName:       "testName2",
			Capacity:          3,
			PriceInCentiUnits: 400,
			PaymentDay:        20,
			CreatedAt:         sql.NullTime{Time: time.Now()},
		},
	}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetSubscriptionsForUserMock.Return(testData, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	subscriptions, err := svc.GetSubscriptionsForUser(ctx, &api.ID{Id: int64(testUserID)})

	assert.Nil(t, err)
	for i, subscription := range subscriptions.Subscriptions {
		testSubscription := testData[i]
		assert.Equal(t, subscription.Id, int64(testSubscription.ID))
		assert.Equal(t, subscription.ChatId, int64(testSubscription.ChatID))
		assert.Equal(t, subscription.ServiceName, testSubscription.ServiceName)
		assert.Equal(t, subscription.Capacity, int64(testSubscription.Capacity))
		assert.Equal(t, subscription.PriceInCentiUnits, int64(testSubscription.PriceInCentiUnits))
		assert.Equal(t, subscription.PaymentDay, int64(testSubscription.PaymentDay))
		assert.Equal(t, subscription.CreatedAt, testSubscription.CreatedAt.Time.Format(time.RFC3339))
	}
}

func TestGetSubscriptionsNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetSubscriptionsForUserMock.Return([]models.Subscription{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.GetSubscriptionsForUser(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
