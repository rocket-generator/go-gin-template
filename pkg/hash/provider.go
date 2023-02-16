package hash

import "golang.org/x/crypto/bcrypt"

type ProviderInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type Provider struct{}

func (p *Provider) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *Provider) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewHashProvider() (ProviderInterface, error) {
	return &Provider{}, nil
}
