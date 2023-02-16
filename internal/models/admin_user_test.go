package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFakeAdminUser(t *testing.T) {
	t.Run("Get Fake Admin User", func(t *testing.T) {
		model, err := GetFakeAdminUser()
		assert.NoError(t, err, "GetFakeAdminUser should not return error")
		assert.NotNil(t, model, "GetFakeAdminUser should return fake admin user")
	})
}
