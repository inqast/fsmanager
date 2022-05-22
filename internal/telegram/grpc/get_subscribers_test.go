package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSubscribers(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testSubscriptionId := 2

	testData := []*api.Subscriber{
		{
			Id:             1,
			UserID:         2,
			SubscriptionID: 3,
			IsPaid:         false,
			IsOwner:        true,
			CreatedAt:      time.Now().Format(time.RFC3339),
		},
		{
			Id:             2,
			UserID:         1,
			SubscriptionID: 3,
			IsPaid:         false,
			IsOwner:        false,
			CreatedAt:      time.Now().Format(time.RFC3339),
		},
	}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetSubscribersForSubscriptionMock.Return(&api.GetSubscribersResponse{Subscribers: testData}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	subscribers, err := svc.GetSubscribers(ctx, testSubscriptionId)

	assert.Nil(t, err)
	for i, subscriber := range subscribers {
		testSubscriber := testData[i]
		assert.Equal(t, subscriber.ID, int(testSubscriber.Id))
		assert.Equal(t, subscriber.UserID, int(testSubscriber.UserID))
		assert.Equal(t, subscriber.SubscriptionID, int(testSubscriber.SubscriptionID))
		assert.Equal(t, subscriber.IsPaid, testSubscriber.IsPaid)
		assert.Equal(t, subscriber.IsOwner, testSubscriber.IsOwner)
		assert.Equal(t, subscriber.CreatedAt.Time.Format(time.RFC3339), testSubscriber.CreatedAt)
	}
}

func TestGetSubscribersNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.GetSubscribersForSubscriptionMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.GetSubscribers(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
