package subscription

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testChatID := 1

	testSubscriptionId := 1
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDay := 28

	testUserId := 1
	testUserTelegramId := 12131
	testUserName := "testName"
	testUserTime := sql.NullTime{Time: time.Now()}

	ctx := context.Background()

	authMock := NewAuthMock(mc)
	authMock.AuthenticateMock.Return(&models.User{
		ID:         testUserId,
		TelegramID: testUserTelegramId,
		Name:       testUserName,
		Pwd:        "",
		CreatedAt:  testUserTime,
	}, nil)
	authMock.AuthenticateMock.Expect(ctx, testUserTelegramId, testUserName)

	sendMock := NewSenderMock(mc)

	grpcMock := NewGrpcClientMock(mc)
	grpcMock.CreateSubscriptionMock.Return(1, nil)
	grpcMock.CreateSubscriptionMock.Expect(ctx, &models.Subscription{
		ChatID:            testChatID,
		ServiceName:       testServiceName,
		Capacity:          testCapacity,
		PriceInCentiUnits: testPriceInCentiUnits,
		PaymentDay:        testPaymentDay,
	})
	grpcMock.CreateSubscriberMock.Return(1, nil)
	grpcMock.CreateSubscriberMock.Expect(ctx, &models.Subscriber{
		UserID:         testUserId,
		SubscriptionID: testSubscriptionId,
		IsPaid:         false,
		IsOwner:        true,
	})
	svc := New(sendMock, grpcMock, authMock)

	id, err := svc.add(
		ctx,
		testChatID,
		testUserTelegramId,
		testUserName,
		[]string{"name=testService", "cap=5", "cost=5", "payday=28"})

	assert.Nil(t, err)
	assert.Equal(t, id.Id, testSubscriptionId)
}
