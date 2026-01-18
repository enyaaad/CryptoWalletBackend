package service

import (
	"context"

	"github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/domain/auth/entity"
)

type AuthService interface {
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.AuthResponse, error)
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.AuthResponse, error)
	ValidateToken(ctx context.Context, tokenString string) (*entity.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*entity.AuthResponse, error)
}
