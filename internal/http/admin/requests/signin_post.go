package requests

// SigninPost ... Object for Signin
type SigninPost struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
