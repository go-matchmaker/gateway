package dto

type UserLoginRequestResponse struct {
	Email        string                 `json:"email"`
	Password     string                 `json:"password"`
	Name         string                 `json:"name"`
	Surname      string                 `json:"surname"`
	PhoneNumber  string                 `json:"phone_number"`
	DepartmentID int                    `json:"department_id"`
	Attributes   map[string]Permissions `json:"attributes"`
}
