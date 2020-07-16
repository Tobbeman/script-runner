package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) registerRS(root *echo.Group) {
	g := root.Group("rs")

	g.GET("/list", h.listScripts)
	g.Any("/:script", h.runScript)
	g.Any("/async/:script", h.asyncRunScript)
	g.Any("/async/:uuid/status", h.asyncStatus)
}

func (h *Handler) listScripts(ctx echo.Context) error {
	if !contains(h.config.Token, extractToken(h.config.ReadTokenHeaders, ctx.Request().Header)) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	res, err := h.runner.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) runScript(ctx echo.Context) error {
	if !contains(h.config.Token, extractToken(h.config.ReadTokenHeaders, ctx.Request().Header)) {
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

func (h *Handler) asyncRunScript(ctx echo.Context) error {
	if !contains(h.config.Token, extractToken(h.config.ReadTokenHeaders, ctx.Request().Header)) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	script := ctx.Param("script")
	if script == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	res, err := h.runner.RunAsync(script, []string{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	uuid := h.store.Store(res)
	ctx.Response().Header().Set("Location", h.config.HrefAddress+"/async/"+uuid+"/status")
	return ctx.String(http.StatusCreated, uuid)
}

func (h *Handler) asyncStatus(ctx echo.Context) error {
	if !contains(h.config.Token, extractToken(h.config.ReadTokenHeaders, ctx.Request().Header)) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	uuid := ctx.Param("uuid")
	if uuid == "" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	found, cmd := h.store.Get(uuid)
	if !found {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	done := cmd.CheckDone()
	var output string
	if done {
		o, err := cmd.Collect()
		if err != nil {
			return echo.NewHTTPError(http.StatusConflict)
		}
		output = o
	}

	res := struct {
		Done   bool   `json:"done"`
		Output string `json:"output"`
	}{
		done,
		output,
	}

	return ctx.JSON(http.StatusOK, res)
}

func extractToken(allowed []string, headers map[string][]string) []string {
	var foundTokens = []string{}
	for key, value := range headers {
		if contains(key, allowed) {
			for _, token := range value {
				foundTokens = append(foundTokens, token)
			}
		}
	}
	return foundTokens
}

func contains(val string, list []string) bool {
	for _, x := range list {
		if x == val {
			return true
		}
	}
	return false
}
