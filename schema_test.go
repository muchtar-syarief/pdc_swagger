package pdc_swagger_test

import (
	"encoding/json"
	"testing"

	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/stretchr/testify/assert"
)

type ApiError struct {
	Message           string            `json:"message,omitempty"`
	ValidationMessage map[string]string `json:"validation_message"`
}

type ApiResponse[T any] struct {
	*ApiError
	Data T `json:"data"`
}

type User struct {
	Email    string `json:"email" fmt:"email" binding:"required,email"`
	Password string `json:"password" fmt:"password" binding:"required,password"`
	Username string `json:"username" binding:"required,gte=6,lte=32"`
}

type ListUser []*User

type UserMapObject struct {
	UserMapper map[string]*User `json:"user_mapper"`
}

func TestBuildSchema(t *testing.T) {

	t.Run("test not pointer", func(t *testing.T) {
		apiError := ApiError{}
		result := pdc_swagger.NewSchema(apiError)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test not pointer with pointer", func(t *testing.T) {
		apiError := &ApiError{}
		result := pdc_swagger.NewSchema(apiError)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test embed struct and generic", func(t *testing.T) {
		response := &ApiResponse[map[string]string]{}
		result := pdc_swagger.NewSchema(response)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test fmt tag", func(t *testing.T) {
		usr := &User{}
		result := pdc_swagger.NewSchema(usr)
		assert.NotEmpty(t, result)

		raw, err := json.MarshalIndent(result, "", "	")
		assert.Nil(t, err)
		t.Log(string(raw))
	})

	t.Run("test map to struct", func(t *testing.T) {
		usr := &UserMapObject{}
		result := pdc_swagger.NewSchema(usr)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test array", func(t *testing.T) {
		data := &ListUser{}
		result := pdc_swagger.NewSchema(data)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})
}
