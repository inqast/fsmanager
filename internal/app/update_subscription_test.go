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
	testOwnerID := 1
	testServiceName := "testService"
	testCapacity := 5
	testPriceInCentiUnits := 500
	testPaymentDate := sql.NullTime{Time: time.Now()}
	testCreatedAt := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateSubscriptionMock.Return(nil)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateSubscription(ctx, &api.Subscription{
		Id:                int64(testId),
		OwnerID:           int64(testOwnerID),
		ServiceName:       testServiceName,
		Capacity:          int64(testCapacity),
		PriceInCentiUnits: int64(testPriceInCentiUnits),
		PaymentDate:       testPaymentDate.Time.Format(time.RFC3339),
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
	_, err := svc.UpdateSubscription(ctx, &api.Subscription{Id: 1, PaymentDate: time.Now().Format(time.RFC3339)})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
