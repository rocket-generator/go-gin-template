package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTokenProvider(t *testing.T) {
	t.Run("Create New Token Provider", func(t *testing.T) {
		got, err := NewTokenProvider([]byte("secret"))
		assert.NoError(t, err, "NewTokenProvider should not return error")
		assert.NotNil(t, got, "NewTokenProvider should return token provider")
	})
}

func TestProvider_GenerateToken(t *testing.T) {
	t.Run("Generate token with GenerateToken", func(t *testing.T) {
		got, _ := NewTokenProvider([]byte("secret"))

		_, err := got.GenerateToken("userID", map[string]interface{}{
			"email": "test@example.com",
		})
		assert.NoError(t, err, "GenerateToken should not return error")
	})
}

func TestProvider_ValidateToken(t *testing.T) {
	t.Run("Validate token with ValidateToken", func(t *testing.T) {
		got, _ := NewTokenProvider([]byte("secret"))
		email := "test@example.com"
		token, _ := got.GenerateToken("userID", map[string]interface{}{
			"email": email,
		})
		check, err := got.ValidateToken(token)
		assert.NoError(t, err, "ValidateToken should not return error")
		assert.NotNil(t, check, "ValidateToken should return token")
		assert.Equal(t, email, check["UserInfo"].(map[string]interface{})["email"])
	})
}
