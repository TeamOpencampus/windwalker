package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"windwalker/models"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	db, err := models.SetupDatabase()
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
}
