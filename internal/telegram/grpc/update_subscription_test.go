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

func TestUpdateSubscription(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testChatId := 1
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDay := 19
	testCreatedAt := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateSubscriptionMock.Return(nil, nil)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateSubscription(ctx, &models.Subscription{
		ID:                testId,
		ChatID:            testChatId,
		ServiceName:       testServiceName,
		Capacity:          testCapacity,
		PriceInCentiUnits: testPriceInCentiUnits,
		PaymentDay:        testPaymentDay,
		CreatedAt:         testCreatedAt,
	})

	assert.Nil(t, err)
}

func TestUpdateSubscriptionNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateSubscriptionMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateSubscription(ctx, &models.Subscription{ID: 1})

	assert.Equal(t, err, ErrNotFound)
}
