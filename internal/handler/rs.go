package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) registerRS(root *echo.Group) {
	g := root.Group("rs")

	g.POST("/:script", h.runScript)
}

func (h *Handler) runScript (ctx echo.Context) error {
	if ! contains(h.config.Token, extractToken(ctx.Request().Header)) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	script := ctx.Param("script")
	if script == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	res, err := h.runner.Run(script, []string{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	return ctx.String(http.StatusOK, res)
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