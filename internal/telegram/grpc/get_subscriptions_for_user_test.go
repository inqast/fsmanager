package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionsForUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testSubscriptionId := 2

	testData := []*api.Subscription{
		{
			Id:                1,
			ChatId:            2,
			ServiceName:       "testName",
			Capacity:          5,
			PriceInCentiUnits: 500,
			PaymentDay:        12,
			CreatedAt:         time.Now().Format(time.RFC3339),
		},
		{
			Id:                2,
			ChatId:            1,
			ServiceName:       "testName2",
			Capacity:          3,
			PriceInCentiUnits: 400,
			PaymentDay:        20,
			CreatedAt:         time.Now().Format(time.RFC3339),
		},
	}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetSubscriptionsForUserMock.Return(&api.GetSubscriptionsResponse{Subscriptions: testData}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	subscribers, err := svc.GetSubscriptionsForUser(ctx, testSubscriptionId)

	assert.Nil(t, err)
	for i, subscription := range subscribers {
		testSubscription := testData[i]
		assert.Equal(t, subscription.ID, int(testSubscription.Id))
		assert.Equal(t, subscription.ChatID, int(testSubscription.ChatId))
		assert.Equal(t, subscription.ServiceName, testSubscription.ServiceName)
		assert.Equal(t, subscription.Capacity, int(testSubscription.Capacity))
		assert.Equal(t, subscription.PriceInCentiUnits, int(testSubscription.PriceInCentiUnits))
		assert.Equal(t, subscription.PaymentDay, int(testSubscription.PaymentDay))
		assert.Equal(t, subscription.CreatedAt.Time.Format(time.RFC3339), testSubscription.CreatedAt)
	}
}

func TestGetSubscriptionsForUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetSubscriptionsForUserMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.GetSubscriptionsForUser(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
