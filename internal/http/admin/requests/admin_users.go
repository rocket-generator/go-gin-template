package requests

// AdminUserCreate ... Object for create admin user
type AdminUserCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AdminUserUpdate ... Object for update admin user
type AdminUserUpdate struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
