package server

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-server/app/config"
	"grpc-server/grpc/paseto"
	auth "grpc-server/grpc/proto"
	"log"
	"net"
	"time"
)

type GrpcServer struct {
	auth.AuthServiceServer // Embed the AuthServiceServer interface
	pasetoMaker            *paseto.Maker
	tokenVerifyMap         map[string]*auth.AuthData // this will be replaced with database or cache in the future
}

func NewGrpcServer(config *config.Config) {
	pasetoMaker := paseto.NewPasetoMaker(config)
	if pasetoMaker == nil {
		panic("failed to create paseto maker")
	}

	listen, err := net.Listen("tcp", *config.Grpc.Url)
	if err != nil {
		panic("failed to listen gRPC server on " + *config.Grpc.Url + ": " + err.Error())
	}

	server := grpc.NewServer([]grpc.ServerOption{}...)
	authServer := &GrpcServer{
		pasetoMaker:    pasetoMaker,
		tokenVerifyMap: make(map[string]*auth.AuthData),
	}
	auth.RegisterAuthServiceServer(server, authServer)
	reflection.Register(server)

	go func() {
		log.Println("Starting gRPC server on " + *config.Grpc.Url)
		err = server.Serve(listen)
		if err != nil {
			panic("failed to serve gRPC server on " + *config.Grpc.Url + ": " + err.Error())
		}
	}()
}

// Override the methods from the AuthServiceServer interface

func (s *GrpcServer) CreateToken(_ context.Context, request *auth.CreateTokenRequest) (*auth.CreateTokenResponse, error) {
	data := request.GetAuthData()
	token := data.GetToken()
	s.tokenVerifyMap[token] = data
	return &auth.CreateTokenResponse{AuthData: data}, nil
}

func (s *GrpcServer) VerifyToken(_ context.Context, request *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {
	token := request.GetToken()
	response := &auth.VerifyTokenResponse{
		Verify: &auth.Verify{AuthData: nil},
	}

	// Check if the token exists in the map
	authData, ok := s.tokenVerifyMap[token]
	if !ok {
		response.Verify.Status = auth.ResponseType_FAILURE
		return response, errors.New("token not found")
	}

	// Verify the token using the paseto maker
	err := s.pasetoMaker.VerifyToken(token)
	if err != nil {
		return nil, errors.New("failed to verify token: " + err.Error())
	}

	// Check if the token is expired
	if authData.ExpiresAt < time.Now().Unix() {
		response.Verify.Status = auth.ResponseType_EXPIRED
		return response, errors.New("token is expired")
	}

	response.Verify.AuthData = authData
	response.Verify.Status = auth.ResponseType_SUCCESS
	return response, nil
}

func (s *GrpcServer) mustEmbedUnimplementedAuthServiceServer() {
	panic("Not implemented yet")
}
