package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"windwalker/models"

	"github.com/stretchr/testify/assert"
)

type AuthResponse struct {
	ID           int    `json:"id"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func TestMain(t *testing.T) {
	db, err := models.SetupDatabase(":memory:")
	assert.NoError(t, err, "Database should setup correctly.")
	r := SetupRouter(db)
	t.Run("HealthCheckHandler", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/v1/health", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, `"all systems operational"`, string(res))
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("RegisterHandler should create account", func(t *testing.T) {
		form, _ := json.Marshal(map[string]string{"email": "someone@example.com", "password": "12345678"})
		req, _ := http.NewRequest(http.MethodPost, "/v1/register", bytes.NewBuffer(form))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var resp AuthResponse

		body, _ := io.ReadAll(w.Body)
		err = json.Unmarshal(body, &resp)
		assert.NoError(t, err)

		assert.NotEmpty(t, resp.Token)
		assert.NotEmpty(t, resp.RefreshToken)
		assert.Equal(t, "user", resp.Role)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("RegisterHandler should not create account with existing email", func(t *testing.T) {
		form, _ := json.Marshal(map[string]string{"email": "someone@example.com", "password": "12345678"})
		req, _ := http.NewRequest(http.MethodPost, "/v1/register", bytes.NewBuffer(form))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var resp ErrorResponse

		body, _ := io.ReadAll(w.Body)
		err = json.Unmarshal(body, &resp)
		assert.NoError(t, err)

		assert.Equal(t, "auth/email-already-registered", resp.Code)
	})

	t.Run("LoginHandler should success", func(t *testing.T) {
		form, _ := json.Marshal(map[string]string{"email": "someone@example.com", "password": "12345678"})
		req, _ := http.NewRequest(http.MethodPost, "/v1/login", bytes.NewBuffer(form))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var resp AuthResponse

		body, _ := io.ReadAll(w.Body)
		err = json.Unmarshal(body, &resp)
		assert.NoError(t, err)

		assert.NotEmpty(t, resp.Token)
		assert.NotEmpty(t, resp.RefreshToken)
		assert.Equal(t, "user", resp.Role)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("LoginHandler should not login with non existant email", func(t *testing.T) {
		form, _ := json.Marshal(map[string]string{"email": "nonexistant@example.com", "password": "12345678"})
		req, _ := http.NewRequest(http.MethodPost, "/v1/login", bytes.NewBuffer(form))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var resp ErrorResponse

		body, _ := io.ReadAll(w.Body)
		err = json.Unmarshal(body, &resp)
		assert.NoError(t, err)

		assert.Equal(t, "auth/account-nonexistent", resp.Code)
	})

	t.Run("LoginHandler should not login with invalid password", func(t *testing.T) {
		form, _ := json.Marshal(map[string]string{"email": "someone@example.com", "password": "invalid_password"})
		req, _ := http.NewRequest(http.MethodPost, "/v1/login", bytes.NewBuffer(form))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var resp ErrorResponse

		body, _ := io.ReadAll(w.Body)
		err = json.Unmarshal(body, &resp)
		assert.NoError(t, err)

		assert.Equal(t, "auth/password-invalid", resp.Code)
	})
}
