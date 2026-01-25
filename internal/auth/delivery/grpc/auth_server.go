package grpc

import (
	"context"

	authProto "github.com/enyaaad/CryptoWalletBackend/api/proto"
	authDomain "github.com/enyaaad/CryptoWalletBackend/internal/auth/domain"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/entity"
	authServiceInterface "github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authGRPCServer struct {
	authProto.UnimplementedAuthServiceServer
	authService authServiceInterface.AuthService
	logger      zerolog.Logger
}

func NewAuthGRPCServer(
	authService authServiceInterface.AuthService,
	logger zerolog.Logger,
) authProto.AuthServiceServer {
	return &authGRPCServer{
		authService: authService,
		logger:      logger,
	}
}

func (s *authGRPCServer) Register(ctx context.Context, req *authProto.RegisterRequest) (*authProto.AuthResponse, error) {
	domainReq := &entity.RegisterRequest{
		Email:    req.GetEmail(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	domainResp, err := s.authService.Register(ctx, domainReq)
	if err != nil {
		return nil, s.mapDomainError(err)
	}

	return s.toProtoAuthResponse(domainResp), nil
}

func (s *authGRPCServer) Login(ctx context.Context, req *authProto.LoginRequest) (*authProto.AuthResponse, error) {
	domainReq := &entity.LoginRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	domainResp, err := s.authService.Login(ctx, domainReq)
	if err != nil {
		return nil, s.mapDomainError(err)
	}

	return s.toProtoAuthResponse(domainResp), nil
}

func (s *authGRPCServer) RefreshToken(ctx context.Context, req *authProto.RefreshTokenRequest) (*authProto.AuthResponse, error) {
	domainResp, err := s.authService.RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, s.mapDomainError(err)
	}

	return s.toProtoAuthResponse(domainResp), nil
}

func (s *authGRPCServer) ValidateToken(ctx context.Context, req *authProto.ValidateTokenRequest) (*authProto.ValidateTokenResponse, error) {
	user, err := s.authService.ValidateToken(ctx, req.GetToken())
	if err != nil {
		if err == authDomain.ErrInvalidToken || err == authDomain.ErrTokenExpired {
			return &authProto.ValidateTokenResponse{
				Valid: false,
			}, nil
		}
		return nil, s.mapDomainError(err)
	}

	return &authProto.ValidateTokenResponse{
		Valid: true,
		User: &authProto.UserInfo{
			Id:       int32(user.ID),
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (s *authGRPCServer) toProtoAuthResponse(domainResp *entity.AuthResponse) *authProto.AuthResponse {
	if domainResp == nil {
		return nil
	}

	protoResp := &authProto.AuthResponse{
		AccessToken:  domainResp.AccessToken,
		RefreshToken: domainResp.RefreshToken,
		TokenType:    domainResp.TokenType,
		ExpiresIn:    int32(domainResp.ExpiresIn),
	}

	if domainResp.User != nil {
		protoResp.User = &authProto.UserInfo{
			Id:       int32(domainResp.User.ID),
			Email:    domainResp.User.Email,
			Username: domainResp.User.Username,
		}
	}

	return protoResp
}

func (s *authGRPCServer) mapDomainError(err error) error {
	switch err {
	case authDomain.ErrUserNotFound:
		return status.Error(codes.NotFound, err.Error())
	case authDomain.ErrInvalidPassword:
		return status.Error(codes.Unauthenticated, err.Error())
	case authDomain.ErrUserAlreadyExists:
		return status.Error(codes.AlreadyExists, err.Error())
	case authDomain.ErrInvalidToken, authDomain.ErrTokenExpired:
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		s.logger.Error().Err(err).Msg("Unexpected error in gRPC handler")
		return status.Error(codes.Internal, "internal server error")
	}
}
