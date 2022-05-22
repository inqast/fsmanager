package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateSubscriptionMock.Return(nil)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateSubscription(ctx, &api.Subscription{
		Id:                int64(testId),
		ChatId:            int64(testChatId),
		ServiceName:       testServiceName,
		Capacity:          int64(testCapacity),
		PriceInCentiUnits: int64(testPriceInCentiUnits),
		PaymentDay:        int64(testPaymentDay),
		CreatedAt:         testCreatedAt.Time.Format(time.RFC3339),
	})

	assert.Nil(t, err)
}

func TestUpdateSubscriptionNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateSubscriptionMock.Return(repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateSubscription(ctx, &api.Subscription{Id: 1, PaymentDay: 19})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
