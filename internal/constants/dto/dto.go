package dto

type AuthRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	XforwardedFor string `json:"x_forwarded_for"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
