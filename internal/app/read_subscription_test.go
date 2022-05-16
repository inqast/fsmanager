package app

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/internal/models"
	"github.com/inqast/fsmanager/internal/repository"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestReadSubscription(t *testing.T) {

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
	mockRepo.ReadSubscriptionMock.Return(models.Subscription{
		ID:                testId,
		OwnerID:           testOwnerID,
		ServiceName:       testServiceName,
		Capacity:          testCapacity,
		PriceInCentiUnits: testPriceInCentiUnits,
		PaymentDate:       testPaymentDate,
		CreatedAt:         testCreatedAt,
	}, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	subscription, err := svc.ReadSubscription(ctx, &api.ID{Id: int64(testId)})

	assert.Nil(t, err)
	assert.Equal(t, subscription.Id, int64(testId))
	assert.Equal(t, subscription.OwnerID, int64(testOwnerID))
	assert.Equal(t, subscription.ServiceName, testServiceName)
	assert.Equal(t, subscription.Capacity, int64(testCapacity))
	assert.Equal(t, subscription.PriceInCentiUnits, int64(testPriceInCentiUnits))
	assert.Equal(t, subscription.PaymentDate, testPaymentDate.Time.Format(time.RFC3339))
	assert.Equal(t, subscription.CreatedAt, testCreatedAt.Time.Format(time.RFC3339))
}

func TestReadSubscriptionNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadSubscriptionMock.Return(models.Subscription{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.ReadSubscription(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
