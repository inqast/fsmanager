package grpc

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSubscription(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1

	clientMock := NewFamilySubClientMock(mc)
	clientMock.DeleteSubscriptionMock.Return(nil, nil)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.DeleteSubscription(ctx, int64(testId))

	assert.Nil(t, err)
}

func TestDeleteSubscriptionNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.DeleteSubscriptionMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.DeleteSubscription(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
