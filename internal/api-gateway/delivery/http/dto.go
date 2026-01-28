package http

// ============== Request DTO ====================

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// ============== Response DTO ===================

type AuthResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	TokenType    string   `json:"token_type"`
	ExpiresIn    int32    `json:"expires_in"`
	User         UserInfo `json:"user"`
}

type UserInfo struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ValidateTokenResponse struct {
	Valid bool      `json:"valid"`
	User  *UserInfo `json:"user,omitempty"`
}
