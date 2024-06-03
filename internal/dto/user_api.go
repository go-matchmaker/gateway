package dto

import (
	"time"
)

// Requests
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

// Responses
type UserLoginResponse struct {
	Data    UserDetail `json:"data"`
	Message string     `json:"message"`
	Status  int        `json:"status"`
}

type UserDetail struct {
	ID            string `json:"id"`
	AccessToken   string `json:"access_token"`
	AccessPublic  string `json:"access_public"`
	RefreshToken  string `json:"refresh_token"`
	RefreshPublic string `json:"refresh_public"`
}

type GetUserResponse struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Surname      string                 `json:"surname"`
	Email        string                 `json:"email"`
	PhoneNumber  string                 `json:"phone_number"`
	Role         string                 `json:"role"`
	DepartmentID string                 `json:"department_id"`
	Attributes   map[string]*Permission `json:"attributes"`
	CreatedAt    time.Time              `json:"created_at"`
}
