package hash

type ProviderMock struct{}

func (p *ProviderMock) HashPassword(password string) (string, error) {
	return "hashed_password", nil
}

func (p *ProviderMock) CheckPasswordHash(password, hash string) bool {
	return true
}
