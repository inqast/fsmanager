package grpc

import "errors"

type Service struct {
	grpcClient  FamilySubClient
	ErrNotFound error
}

var ErrNotFound = errors.New("not found")

func New(client FamilySubClient) *Service {
	return &Service{
		grpcClient: client,
	}
}
