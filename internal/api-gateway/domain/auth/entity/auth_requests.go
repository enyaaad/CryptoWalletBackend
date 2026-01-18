package entity

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Username string `json:"username" validate:"required,min=3,max=50" example:"johndoe"`
	Password string `json:"password" validate:"required,min=8" example:"password123"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
}
