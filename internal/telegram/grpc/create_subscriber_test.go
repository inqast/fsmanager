package grpc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/inqast/fsmanager/internal/models"

	"github.com/gojuno/minimock/v3"
	"github.com/inqast/fsmanager/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubscriber(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1
	testUserID := 2
	testSubscriptionID := 3
	testTime := sql.NullTime{Time: time.Now()}

	clientMock := NewFamilySubClientMock(mc)
	clientMock.CreateSubscriberMock.Return(&api.ID{Id: int64(testId)}, nil)
	svc := New(clientMock)

	ctx := context.Background()
	id, err := svc.CreateSubscriber(ctx, &models.Subscriber{
		ID:             testId,
		UserID:         testUserID,
		SubscriptionID: testSubscriptionID,
		IsPaid:         false,
		IsOwner:        true,
		CreatedAt:      testTime,
	})

	assert.Nil(t, err)
	assert.Equal(t, id, testId)
}
