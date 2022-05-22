package subscription

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/inqast/fsmanager/internal/telegram/message"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestPay(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testChatID := 1

	testSubscriptionId := 1
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDay := 28
	testSubscriptionCreateAt := sql.NullTime{Time: time.Now()}

	testUserId := 2
	testUserTelegramId := 12131
	testUserName := "testName"
	testUserTime := sql.NullTime{Time: time.Now()}

	testSubscriberId := 1
	testSubscriberCreateAt := sql.NullTime{Time: time.Now()}

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
	grpcMock.ReadSubscriptionMock.Return(&models.Subscription{
		ID:                testSubscriptionId,
		ChatID:            testChatID,
		ServiceName:       testServiceName,
		Capacity:          testCapacity,
		PriceInCentiUnits: testPriceInCentiUnits,
		PaymentDay:        testPaymentDay,
		CreatedAt:         testSubscriptionCreateAt,
	}, nil)
	grpcMock.ReadSubscriptionMock.Expect(ctx, testSubscriptionId)

	grpcMock.GetSubscribersMock.Return([]*models.Subscriber{
		{
			ID:             testSubscriberId,
			UserID:         testUserId,
			SubscriptionID: testSubscriptionId,
			IsPaid:         false,
			IsOwner:        true,
			CreatedAt:      testSubscriberCreateAt,
		},
	}, nil)
	grpcMock.GetSubscribersMock.Expect(ctx, testSubscriptionId)

	grpcMock.UpdateSubscriberMock.Return(true, nil)
	grpcMock.UpdateSubscriberMock.Expect(ctx, &models.Subscriber{
		ID:             testSubscriberId,
		UserID:         testUserId,
		SubscriptionID: testSubscriptionId,
		IsPaid:         true,
		IsOwner:        true,
		CreatedAt:      testSubscriberCreateAt,
	})

	svc := New(sendMock, grpcMock, authMock)

	result, err := svc.pay(
		ctx,
		testChatID,
		testUserTelegramId,
		testUserName,
		[]string{"1"})

	assert.Nil(t, err)
	assert.Equal(t, &message.Success{}, result)
}
