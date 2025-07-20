package repository

import (
	"grpc-server/app/config"
	"grpc-server/grpc/client"
)

type Repository struct {
	config     *config.Config
	grpcClient *client.GrpcClient
}

func NewRepository(config *config.Config, grpcClient *client.GrpcClient) (*Repository, error) {
	return &Repository{
		config:     config,
		grpcClient: grpcClient,
	}, nil
}

func (r Repository) CreateAuth(name string) (interface{}, error) {
	// Implement the logic to create an auth entity using the gRPC client
	// This is a placeholder implementation
	authData, err := r.grpcClient.CreateToken(name)
	if err != nil {
		return nil, err
	}
	return authData, nil
}
