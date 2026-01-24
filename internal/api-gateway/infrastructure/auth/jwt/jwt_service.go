package jwt_service

import (
	"strconv"
	"time"

	"github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/domain/auth"
	"github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/domain/auth/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtService interface {
	GenerateAccessToken(user *entity.User) (string, error)
	GenereteRefreshToken(user *entity.User) (string, error)
	ValidateToken(tokenString string) (*entity.JWTClaims, error)
}
type jwt_service struct {
	secretKey     []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService(secretKey string, accessExpiry, refreshExpiry time.Duration) JwtService {
	return &jwt_service{
		secretKey:     []byte(secretKey),
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (jwts *jwt_service) GenerateAccessToken(user *entity.User) (string, error) {
	now := time.Now()

	claims := entity.JWTClaims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(jwts.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "cryptowallet-api",
			Subject:   strconv.Itoa(user.ID),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwts.secretKey)
}

func (jwts *jwt_service) GenereteRefreshToken(user *entity.User) (string, error) {
	now := time.Now()

	claims := entity.JWTClaims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(jwts.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "cryptowallet-api",
			Subject:   strconv.Itoa(user.ID),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwts.secretKey)
}

func (jwts *jwt_service) ValidateToken(tokenString string) (*entity.JWTClaims, error) {
	claims := &entity.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return jwts.secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, auth.ErrTokenExpired
		}

		if err == jwt.ErrSignatureInvalid {
			return nil, auth.ErrInvalidToken
		}
		return nil, auth.ErrInvalidToken
	}

	if !token.Valid {
		return nil, auth.ErrInvalidToken
	}

	if claims.Issuer != "cryptowallet-api" {
		return nil, auth.ErrInvalidToken
	}

	return claims, nil
}
