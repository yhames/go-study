package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-server/app/config"
	"grpc-server/grpc/paseto"
	auth "grpc-server/grpc/proto"
	"time"
)

type GrpcClient struct {
	conn        *grpc.ClientConn
	client      auth.AuthServiceClient
	pasetoMaker *paseto.Maker
}

func NewGrpcClient(config *config.Config) (*GrpcClient, error) {
	g := new(GrpcClient)

	conn, err := grpc.NewClient(*config.Grpc.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	g.conn = conn
	g.client = auth.NewAuthServiceClient(conn)
	g.pasetoMaker = paseto.NewPasetoMaker(config)
	return g, nil
}

func (c *GrpcClient) CreateToken(name string) (*auth.AuthData, error) {
	now := time.Now()
	expiresAt := now.Add(30 * time.Minute) // Default expiration time
	a := &auth.AuthData{
		Name:      name,
		CreatedAt: now.Unix(),
		ExpiresAt: expiresAt.Unix(),
	}

	token, err := c.pasetoMaker.CreateToken(a)
	if err != nil {
		return nil, err
	}
	a.Token = token

	createToken, err := c.client.CreateToken(context.Background(), &auth.CreateTokenRequest{AuthData: a})
	if err != nil {
		return nil, err
	}
	return createToken.AuthData, nil
}

func (c *GrpcClient) VerifyToken(token string) (*auth.VerifyTokenResponse, error) {
	if err := c.pasetoMaker.VerifyToken(token); err != nil {
		return nil, err
	}
	return c.client.VerifyToken(context.Background(), &auth.VerifyTokenRequest{Token: token})
}
