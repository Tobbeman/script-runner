package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunScript(t *testing.T) {
	// Setup
	token := "lol"
	h := New(
		*runner.New("."),
		config.Config{
			Token: token,
		})
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-Gitlab-Token", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.runScript(c)) {
	}

	req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-Gitlab-Token", "token")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	// Assertions
	if assert.Error(t, h.runScript(c)) {
	}
}
