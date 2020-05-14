package handler

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
)

type Handler struct {
	runner *runner.Runner
	config *config.Config
}

func New(r *runner.Runner, c *config.Config) *Handler {
	return &Handler{r, c}
}

func (h *Handler) Register(g *echo.Group) {
	h.registerRS(g)
}
