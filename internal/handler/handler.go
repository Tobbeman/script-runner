package handler

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
)

type Handler struct {
	runner *runner.Runner
	store  *runner.RCmdStore
	config *config.Config
}

func New(r *runner.Runner, s *runner.RCmdStore, c *config.Config) *Handler {
	h := &Handler{
		r,
		s,
		c}
	return h
}

func (h *Handler) Register(g *echo.Group) {
	h.registerRS(g)
}
