package service

import (
	"context"

	auth "github.com/enyaaad/CryptoWalletBackend/internal/auth/domain"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/entity"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/repository"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/service"
	jwt_service "github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/jwt"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/password"
)

type authService struct {
	userRepo       repository.UserRepository
	passwordHasher password.Hasher
	jwtService     jwt_service.JwtService
}

func NewAuthService(
	userRepo repository.UserRepository,
	passwordHasher password.Hasher,
	jwtService jwt_service.JwtService,
) service.AuthService {
	return &authService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		jwtService:     jwtService,
	}
}

func (s *authService) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.AuthResponse, error) {
	_, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		return nil, auth.ErrUserAlreadyExists
	}
	if err != auth.ErrUserNotFound {
		return nil, err
	}

	hashedPassword, err := s.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtService.GenereteRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    15 * 60, // 15 минут в секундах
		User: &entity.UserInfo{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (s *authService) Login(ctx context.Context, req *entity.LoginRequest) (*entity.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, auth.ErrInvalidPassword
	}

	if !s.passwordHasher.Verify(user.Password, req.Password) {
		return nil, auth.ErrInvalidPassword
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtService.GenereteRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    15 * 60,
		User: &entity.UserInfo{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (authS *authService) ValidateToken(ctx context.Context, tokenString string) (*entity.User, error) {
	claims, err := authS.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	user, err := authS.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (authS *authService) RefreshToken(ctx context.Context, refreshToken string) (*entity.AuthResponse, error) {
	claims, err := authS.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := authS.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := authS.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := authS.jwtService.GenereteRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    15 * 60,
		User: &entity.UserInfo{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}
