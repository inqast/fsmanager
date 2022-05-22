package grpc

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testId := 1

	clientMock := NewFamilySubClientMock(mc)
	clientMock.DeleteUserMock.Return(nil, nil)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.DeleteUser(ctx, int64(testId))

	assert.Nil(t, err)
}

func TestDeleteUserNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	clientMock := NewFamilySubClientMock(mc)
	clientMock.DeleteUserMock.Return(nil, ErrNotFound)
	svc := New(clientMock)

	ctx := context.Background()
	_, err := svc.DeleteUser(ctx, 1)

	assert.Equal(t, err, ErrNotFound)
}
