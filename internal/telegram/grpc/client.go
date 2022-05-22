package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/inqast/fsmanager/pkg/api"
	"google.golang.org/grpc"
)

type FamilySubClient interface {
	CreateUser(context.Context, *api.User, ...grpc.CallOption) (*api.ID, error)
	ReadUser(context.Context, *api.ID, ...grpc.CallOption) (*api.User, error)
	UpdateUser(context.Context, *api.User, ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(context.Context, *api.ID, ...grpc.CallOption) (*empty.Empty, error)
	GetSubscriptionsForUser(context.Context, *api.ID, ...grpc.CallOption) (*api.GetSubscriptionsResponse, error)
	GetUserByTelegramID(context.Context, *api.ID, ...grpc.CallOption) (*api.User, error)
	GetUsersByIDs(context.Context, *api.GetUsersByIDsRequest, ...grpc.CallOption) (*api.GetUsersByIDsResponse, error)
	CreateSubscription(context.Context, *api.Subscription, ...grpc.CallOption) (*api.ID, error)
	ReadSubscription(context.Context, *api.ID, ...grpc.CallOption) (*api.Subscription, error)
	UpdateSubscription(context.Context, *api.Subscription, ...grpc.CallOption) (*empty.Empty, error)
	DeleteSubscription(context.Context, *api.ID, ...grpc.CallOption) (*empty.Empty, error)
	GetSubscribersForSubscription(context.Context, *api.ID, ...grpc.CallOption) (*api.GetSubscribersResponse, error)
	CreateSubscriber(context.Context, *api.Subscriber, ...grpc.CallOption) (*api.ID, error)
	ReadSubscriber(context.Context, *api.ID, ...grpc.CallOption) (*api.Subscriber, error)
	UpdateSubscriber(context.Context, *api.Subscriber, ...grpc.CallOption) (*empty.Empty, error)
	DeleteSubscriber(context.Context, *api.ID, ...grpc.CallOption) (*empty.Empty, error)
}
