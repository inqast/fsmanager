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

func TestUpdateSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateSubscriberMock.Return(nil, nil)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateSubscriber(ctx, &models.Subscriber{
		ID:             testId,
		UserID:         testUserID,
		SubscriptionID: testSubscriptionID,
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime,
	})

	assert.Nil(t, err)
}

func TestUpdateSubscriberNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.UpdateSubscriberMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.UpdateSubscriber(ctx, &models.Subscriber{ID: 1})

	assert.Equal(t, err, ErrNotFound)
}
