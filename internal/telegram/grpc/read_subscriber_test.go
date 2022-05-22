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

func TestReadSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadSubscriberMock.Return(&api.Subscriber{
		Id:             int64(testId),
		UserID:         int64(testUserID),
		SubscriptionID: int64(testSubscriptionID),
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime.Time.Format(time.RFC3339),
	}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	subscriber, err := svc.ReadSubscriber(ctx, testId)

	assert.Nil(t, err)
	assert.Equal(t, subscriber.ID, testId)
	assert.Equal(t, subscriber.UserID, testUserID)
	assert.Equal(t, subscriber.SubscriptionID, testSubscriptionID)
	assert.Equal(t, subscriber.IsPaid, false)
	assert.Equal(t, subscriber.IsOwner, true)
	assert.Equal(t, subscriber.CreatedAt.Time.Format(time.RFC3339), testTime.Time.Format(time.RFC3339))
}

func TestReadSubscriberNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadSubscriberMock.Return(&api.Subscriber{}, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.ReadSubscriber(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
