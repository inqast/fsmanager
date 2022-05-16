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

func TestReadSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadSubscriberMock.Return(models.Subscriber{
		ID:             testId,
		UserID:         testUserID,
		SubscriptionID: testSubscriptionID,
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime,
	}, nil)
	svc := New(mockRepo)

	ctx := context.Background()
	subscriber, err := svc.ReadSubscriber(ctx, &api.ID{Id: int64(testId)})

	assert.Nil(t, err)
	assert.Equal(t, subscriber.Id, int64(testId))
	assert.Equal(t, subscriber.UserID, int64(testUserID))
	assert.Equal(t, subscriber.SubscriptionID, int64(testSubscriptionID))
	assert.Equal(t, subscriber.IsPaid, false)
	assert.Equal(t, subscriber.IsOwner, true)
	assert.Equal(t, subscriber.CreatedAt, testTime.Time.Format(time.RFC3339))
}

func TestReadSubscriberNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadSubscriberMock.Return(models.Subscriber{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.ReadSubscriber(ctx, &api.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
