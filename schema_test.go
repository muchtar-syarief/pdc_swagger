package pdc_swagger_test

import (
	"testing"

	"github.com/muchtar-syarief/pdc_swagger"
	"github.com/stretchr/testify/assert"
)

type ApiError struct {
	Message           string            `json:"message"`
	ValidationMessage map[string]string `json:"validation_message"`
}

type ApiResponse[T any] struct {
	*ApiError
	Data T `json:"data"`
}

type User struct {
	Email    string `json:"email" fmt:"email"`
	Password string `json:"password" fmt:"password"`
}

func TestBuildSchema(t *testing.T) {
	t.Run("test not pointer", func(t *testing.T) {
		apiError := ApiError{}
		result, err := pdc_swagger.NewSchema(apiError)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test not pointer with pointer", func(t *testing.T) {
		apiError := &ApiError{}
		result, err := pdc_swagger.NewSchema(apiError)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test embed struct and generic", func(t *testing.T) {
		response := &ApiResponse[map[string]string]{}
		result, err := pdc_swagger.NewSchema(response)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})

	t.Run("test fmt tag", func(t *testing.T) {
		usr := &User{}
		result, err := pdc_swagger.NewSchema(usr)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)

		// raw, err := json.MarshalIndent(result, "", "	")
		// assert.Nil(t, err)
		// t.Log(string(raw))
	})
}
