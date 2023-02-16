package token

type ProviderMock struct {
}

func (p *ProviderMock) GenerateToken(sub string, userInfo interface{}) (string, error) {
	return "random_string", nil
}

func (p *ProviderMock) ValidateToken(tokenString string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"UserInfo": map[string]interface{}{
			"id":    "random_string",
			"email": "random",
		},
	}, nil
}
