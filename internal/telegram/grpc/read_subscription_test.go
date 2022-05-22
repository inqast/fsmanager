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

func TestReadSubscription(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testChatID := 2
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDay := 19
	testCreatedAt := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadSubscriptionMock.Return(&api.Subscription{
		Id:                int64(testId),
		ChatId:            int64(testChatID),
		ServiceName:       testServiceName,
		Capacity:          int64(testCapacity),
		PriceInCentiUnits: int64(testPriceInCentiUnits),
		PaymentDay:        int64(testPaymentDay),
		CreatedAt:         testCreatedAt.Time.Format(time.RFC3339),
	}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	subscription, err := svc.ReadSubscription(ctx, testId)

	assert.Nil(t, err)
	assert.Equal(t, subscription.ID, testId)
	assert.Equal(t, subscription.ChatID, testChatID)
	assert.Equal(t, subscription.ServiceName, testServiceName)
	assert.Equal(t, subscription.Capacity, testCapacity)
	assert.Equal(t, subscription.PriceInCentiUnits, testPriceInCentiUnits)
	assert.Equal(t, subscription.PaymentDay, testPaymentDay)
	assert.Equal(t, subscription.CreatedAt.Time.Format(time.RFC3339), testCreatedAt.Time.Format(time.RFC3339))
}

func TestReadSubscriptionNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.ReadSubscriptionMock.Return(&api.Subscription{}, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.ReadSubscription(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
