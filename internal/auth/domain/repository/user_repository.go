package repository

import (
	"context"

	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
}
