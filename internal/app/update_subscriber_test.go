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

func TestUpdateSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateSubscriberMock.Return(nil)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateSubscriber(ctx, &api.Subscriber{
		Id:             int64(testId),
		UserID:         int64(testUserID),
		SubscriptionID: int64(testSubscriptionID),
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime.Time.Format(time.RFC3339),
	})

	assert.Nil(t, err)
}

func TestUpdateSubscriberNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateSubscriberMock.Return(repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.UpdateSubscriber(ctx, &api.Subscriber{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
