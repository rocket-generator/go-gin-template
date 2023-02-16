package responses

import "github.com/takaaki-mizuno/go-gin-template/internal/models"

// Token ...
type Token struct {
	Token string `json:"token"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// NewToken ... Create new User response
func NewToken(token string, adminUser models.AdminUser) *Token {
	response := &Token{
		Token: token,
		ID:    adminUser.ID.String(),
		Name:  adminUser.Name,
	}
	return response
}
