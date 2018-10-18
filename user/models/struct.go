package user

type User struct {
	ID       string   `json:"id,omitempty"`
	Customer Customer `json:"customer,omitempty"`
	Password string   `json:"password,omitempty"`
}

type Customer struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UpdatePassword struct {
	Password string   `json:"password,omitempty"`
	NewPassword string   `json:"newPassword,omitempty"`
}
