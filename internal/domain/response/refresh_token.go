package response

type RefreshTokenResponse struct {
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
