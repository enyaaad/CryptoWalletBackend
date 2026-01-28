package http

import (
	"github.com/gin-gonic/gin"

	authProto "github.com/enyaaad/CryptoWalletBackend/api/proto"
	authClient "github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/infrastructure/grpc"
)

type AuthHandler struct {
	client *authClient.AuthClient
}

func NewAuthHandler(authClient *authClient.AuthClient) *AuthHandler {
	return &AuthHandler{client: authClient}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	protoReq := &authProto.RegisterRequest{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	resp, err := h.client.Register(ctx.Request.Context(), protoReq)
	if err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	ctx.JSON(200, AuthResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		TokenType:    resp.TokenType,
		ExpiresIn:    resp.ExpiresIn,
		User: UserInfo{
			ID:       resp.User.Id,
			Email:    resp.User.Email,
			Username: resp.User.Username,
		},
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	protoReq := &authProto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := h.client.Login(ctx.Request.Context(), protoReq)
	if err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	ctx.JSON(200, AuthResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		TokenType:    resp.TokenType,
		ExpiresIn:    resp.ExpiresIn,
		User: UserInfo{
			ID:       resp.User.Id,
			Email:    resp.User.Email,
			Username: resp.User.Username,
		},
	})
}

func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var req RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	protoReq := &authProto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	resp, err := h.client.RefreshToken(ctx.Request.Context(), protoReq)
	if err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	ctx.JSON(200, AuthResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		TokenType:    resp.TokenType,
		ExpiresIn:    resp.ExpiresIn,
		User: UserInfo{
			ID:       resp.User.Id,
			Email:    resp.User.Email,
			Username: resp.User.Username,
		},
	})
}

func (h *AuthHandler) ValidateToken(ctx *gin.Context) {
	var req ValidateTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	protoReq := &authProto.ValidateTokenRequest{
		Token: req.Token,
	}

	resp, err := h.client.ValidateToken(ctx.Request.Context(), protoReq)
	if err != nil {
		code, response := gRPCtoHTTPErr(err)
		ctx.JSON(code, response)
		return
	}

	ctx.JSON(200, ValidateTokenResponse{
		Valid: resp.Valid,
		User: &UserInfo{
			ID:       resp.User.Id,
			Email:    resp.User.Email,
			Username: resp.User.Username,
		},
	})
}
