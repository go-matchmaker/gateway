package dto

type UserLoginResponse struct {
	Token     string `json:"token"`
	PublicKey string `json:"public_key"`
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

type Permission struct {
	View        bool `json:"view"`
	Search      bool `json:"search"`
	Detail      bool `json:"detail"`
	Add         bool `json:"add"`
	Update      bool `json:"update"`
	Delete      bool `json:"delete"`
	Export      bool `json:"export"`
	Import      bool `json:"import"`
	CanSeePrice bool `json:"can_see_price"`
}
