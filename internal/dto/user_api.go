package dto

// Requests
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type CreateUserRequest struct {
	User         User                  `json:"user"`
	Permissions  map[string]Permission `json:"permissions"`
	DepartmentID string                `json:"department_id"`
}

// Responses
type UserLoginResponse struct {
	Data    UserLoginData `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}
type UserLoginData struct {
	Token     string `json:"token"`
	PublicKey string `json:"public_key"`
	User      User   `json:"user"`
}
type AuthMiddlewareResponse struct {
	Email        string                `json:"email"`
	Password     string                `json:"password"`
	Name         string                `json:"name"`
	Surname      string                `json:"surname"`
	PhoneNumber  string                `json:"phone_number"`
	DepartmentID int                   `json:"department_id"`
	Attributes   map[string]Permission `json:"attributes"`
}
