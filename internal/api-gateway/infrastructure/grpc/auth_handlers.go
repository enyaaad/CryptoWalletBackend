package http

import (
	"context"
	"time"

	authProto "github.com/enyaaad/CryptoWalletBackend/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client authProto.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthClient(addr string) (*AuthClient, error) {
	// tls for prod, insecure for dev
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	return &AuthClient{
		client: authProto.NewAuthServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AuthClient) Register(ctx context.Context, req *authProto.RegisterRequest) (*authProto.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.client.Register(ctx, req)
}

func (c *AuthClient) Login(ctx context.Context, req *authProto.LoginRequest) (*authProto.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.client.Login(ctx, req)
}

func (c *AuthClient) RefreshToken(ctx context.Context, req *authProto.RefreshTokenRequest) (*authProto.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.client.RefreshToken(ctx, req)
}

func (c *AuthClient) ValidateToken(ctx context.Context, req *authProto.ValidateTokenRequest) (*authProto.ValidateTokenResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return c.client.ValidateToken(ctx, req)
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}
