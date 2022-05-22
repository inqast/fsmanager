// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api.proto

package api

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd        string `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	TelegramId int64  `protobuf:"varint,4,opt,name=telegram_id,json=telegramId,proto3" json:"telegram_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *User) GetTelegramId() int64 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetUsersByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetUsersByIDsRequest) Reset() {
	*x = GetUsersByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersByIDsRequest) ProtoMessage() {}

func (x *GetUsersByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetUsersByIDsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersByIDsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetUsersByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersByIDsResponse) Reset() {
	*x = GetUsersByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersByIDsResponse) ProtoMessage() {}

func (x *GetUsersByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetUsersByIDsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsersByIDsResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatId            int64  `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	ServiceName       string `protobuf:"bytes,3,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	Capacity          int64  `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	PriceInCentiUnits int64  `protobuf:"varint,5,opt,name=priceInCentiUnits,proto3" json:"priceInCentiUnits,omitempty"`
	PaymentDay        int64  `protobuf:"varint,6,opt,name=paymentDay,proto3" json:"paymentDay,omitempty"`
	CreatedAt         string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Subscription) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subscription) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *Subscription) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Subscription) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Subscription) GetPriceInCentiUnits() int64 {
	if x != nil {
		return x.PriceInCentiUnits
	}
	return 0
}

func (x *Subscription) GetPaymentDay() int64 {
	if x != nil {
		return x.PaymentDay
	}
	return 0
}

func (x *Subscription) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *GetSubscriptionsResponse) Reset() {
	*x = GetSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsResponse) ProtoMessage() {}

func (x *GetSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type Subscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID         int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	SubscriptionID int64  `protobuf:"varint,3,opt,name=subscriptionID,proto3" json:"subscriptionID,omitempty"`
	IsPaid         bool   `protobuf:"varint,4,opt,name=isPaid,proto3" json:"isPaid,omitempty"`
	IsOwner        bool   `protobuf:"varint,5,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	CreatedAt      string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Subscriber) Reset() {
	*x = Subscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriber) ProtoMessage() {}

func (x *Subscriber) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscriber.ProtoReflect.Descriptor instead.
func (*Subscriber) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *Subscriber) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subscriber) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Subscriber) GetSubscriptionID() int64 {
	if x != nil {
		return x.SubscriptionID
	}
	return 0
}

func (x *Subscriber) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *Subscriber) GetIsOwner() bool {
	if x != nil {
		return x.IsOwner
	}
	return false
}

func (x *Subscriber) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetSubscribersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscribers []*Subscriber `protobuf:"bytes,1,rep,name=subscribers,proto3" json:"subscribers,omitempty"`
}

func (x *GetSubscribersResponse) Reset() {
	*x = GetSubscribersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscribersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscribersResponse) ProtoMessage() {}

func (x *GetSubscribersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscribersResponse.ProtoReflect.Descriptor instead.
func (*GetSubscribersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetSubscribersResponse) GetSubscribers() []*Subscriber {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x7b, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x28,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x69, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x43, 0x65,
	0x6e, 0x74, 0x69, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x53, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x50, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x32, 0xa4, 0x0a, 0x0a, 0x09, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x53, 0x75, 0x62, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x07,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x35, 0x0a, 0x08,
	0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49,
	0x44, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x49, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x1a, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x44,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x07, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x49, 0x44, 0x12, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x07, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x4d, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x61, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x1a, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49,
	0x44, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x50, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x12, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x71,
	0x61, 0x73, 0x74, 0x2f, 0x66, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(*ID)(nil),                       // 0: api.ID
	(*User)(nil),                     // 1: api.User
	(*GetUsersByIDsRequest)(nil),     // 2: api.GetUsersByIDsRequest
	(*GetUsersByIDsResponse)(nil),    // 3: api.GetUsersByIDsResponse
	(*Subscription)(nil),             // 4: api.Subscription
	(*GetSubscriptionsResponse)(nil), // 5: api.GetSubscriptionsResponse
	(*Subscriber)(nil),               // 6: api.Subscriber
	(*GetSubscribersResponse)(nil),   // 7: api.GetSubscribersResponse
	(*empty.Empty)(nil),              // 8: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	1,  // 0: api.GetUsersByIDsResponse.users:type_name -> api.User
	4,  // 1: api.GetSubscriptionsResponse.subscriptions:type_name -> api.Subscription
	6,  // 2: api.GetSubscribersResponse.subscribers:type_name -> api.Subscriber
	1,  // 3: api.FamilySub.CreateUser:input_type -> api.User
	0,  // 4: api.FamilySub.ReadUser:input_type -> api.ID
	1,  // 5: api.FamilySub.UpdateUser:input_type -> api.User
	0,  // 6: api.FamilySub.DeleteUser:input_type -> api.ID
	0,  // 7: api.FamilySub.GetSubscriptionsForUser:input_type -> api.ID
	0,  // 8: api.FamilySub.GetUserByTelegramID:input_type -> api.ID
	2,  // 9: api.FamilySub.GetUsersByIDs:input_type -> api.GetUsersByIDsRequest
	4,  // 10: api.FamilySub.CreateSubscription:input_type -> api.Subscription
	0,  // 11: api.FamilySub.ReadSubscription:input_type -> api.ID
	4,  // 12: api.FamilySub.UpdateSubscription:input_type -> api.Subscription
	0,  // 13: api.FamilySub.DeleteSubscription:input_type -> api.ID
	0,  // 14: api.FamilySub.GetSubscribersForSubscription:input_type -> api.ID
	6,  // 15: api.FamilySub.CreateSubscriber:input_type -> api.Subscriber
	0,  // 16: api.FamilySub.ReadSubscriber:input_type -> api.ID
	6,  // 17: api.FamilySub.UpdateSubscriber:input_type -> api.Subscriber
	0,  // 18: api.FamilySub.DeleteSubscriber:input_type -> api.ID
	0,  // 19: api.FamilySub.CreateUser:output_type -> api.ID
	1,  // 20: api.FamilySub.ReadUser:output_type -> api.User
	8,  // 21: api.FamilySub.UpdateUser:output_type -> google.protobuf.Empty
	8,  // 22: api.FamilySub.DeleteUser:output_type -> google.protobuf.Empty
	5,  // 23: api.FamilySub.GetSubscriptionsForUser:output_type -> api.GetSubscriptionsResponse
	1,  // 24: api.FamilySub.GetUserByTelegramID:output_type -> api.User
	3,  // 25: api.FamilySub.GetUsersByIDs:output_type -> api.GetUsersByIDsResponse
	0,  // 26: api.FamilySub.CreateSubscription:output_type -> api.ID
	4,  // 27: api.FamilySub.ReadSubscription:output_type -> api.Subscription
	8,  // 28: api.FamilySub.UpdateSubscription:output_type -> google.protobuf.Empty
	8,  // 29: api.FamilySub.DeleteSubscription:output_type -> google.protobuf.Empty
	7,  // 30: api.FamilySub.GetSubscribersForSubscription:output_type -> api.GetSubscribersResponse
	0,  // 31: api.FamilySub.CreateSubscriber:output_type -> api.ID
	6,  // 32: api.FamilySub.ReadSubscriber:output_type -> api.Subscriber
	8,  // 33: api.FamilySub.UpdateSubscriber:output_type -> google.protobuf.Empty
	8,  // 34: api.FamilySub.DeleteSubscriber:output_type -> google.protobuf.Empty
	19, // [19:35] is the sub-list for method output_type
	3,  // [3:19] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersByIDsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersByIDsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscriber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscribersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
