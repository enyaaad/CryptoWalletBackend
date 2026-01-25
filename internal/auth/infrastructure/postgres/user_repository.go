package postgres

import (
	"context"
	"database/sql"
	"time"

	authDomain "github.com/enyaaad/CryptoWalletBackend/internal/auth/domain"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/entity"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
	INSERT INTO users (email, username, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Username, user.Password, now, now).Scan(&user.ID)

	return err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, username, password_hash, created_at, updated_at
	FROM users WHERE email = $1`
	user := &entity.User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, authDomain.ErrUserNotFound
	}

	return user, err
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, email, username, password_hash, created_at, updated_at FROM users WHERE id = $1`

	user := &entity.User{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, authDomain.ErrUserNotFound
	}

	return user, err
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
        UPDATE users
        SET email = $1, username = $2, password_hash = $3, updated_at = $4
        WHERE id = $5
        RETURNING id, email, username, password_hash, created_at, updated_at`

	user.UpdatedAt = time.Now()

	updatedUser := &entity.User{}
	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Username,
		user.Password,
		user.UpdatedAt,
		user.ID).Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.Username,
		&updatedUser.Password,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, authDomain.ErrUserNotFound
	}

	return updatedUser, err
}
