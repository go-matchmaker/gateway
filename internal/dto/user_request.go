package dto

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type UserAuthRequest struct {
	Token     string `json:"token"`
	PublicKey string `json:"public_key"`
}
