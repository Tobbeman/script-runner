package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) registerRunner(root echo.Group) {
	root.POST("", h.runScript)
}

func (h *Handler) runScript (ctx echo.Context) error {
	if ! contains(h.config.Token, extractToken(ctx.Request().Header)) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return nil
}

func extractToken(headers map[string][]string) []string {
	for key, value := range headers {
		if key == "X-Gitlab-Token" {
			return value
		}
	}
	return []string{}
}

func contains(val string, list []string) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}