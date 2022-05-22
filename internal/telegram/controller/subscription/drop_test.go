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

func TestDrop(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testChatID := 1

	testSubscriptionId := 1
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDay := 28
	testSubscriptionCreateAt := sql.NullTime{Time: time.Now()}

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
	grpcMock.ReadSubscriptionMock.Return(&models.Subscription{
		ID:                testSubscriptionId,
		ChatID:            testChatID,
		ServiceName:       testServiceName,
		Capacity:          testCapacity,
		PriceInCentiUnits: testPriceInCentiUnits,
		PaymentDay:        testPaymentDay,
		CreatedAt:         testSubscriptionCreateAt,
	}, nil)
	grpcMock.ReadSubscriptionMock.Expect(ctx, testUserId)

	grpcMock.GetSubscribersMock.Return([]*models.Subscriber{
		{
			ID:             1,
			UserID:         testUserId,
			SubscriptionID: 3,
			IsPaid:         false,
			IsOwner:        true,
			CreatedAt:      sql.NullTime{Time: time.Now()},
		},
	}, nil)
	grpcMock.GetSubscribersMock.Expect(ctx, testSubscriptionId)

	grpcMock.DeleteSubscriptionMock.Return(true, nil)
	grpcMock.DeleteSubscriptionMock.Expect(ctx, int64(testSubscriptionId))

	svc := New(sendMock, grpcMock, authMock)

	result, err := svc.drop(
		ctx,
		testChatID,
		testUserTelegramId,
		testUserName,
		[]string{"1"})

	assert.Nil(t, err)
	assert.Equal(t, &message.Success{}, result)
}
