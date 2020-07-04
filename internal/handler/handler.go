package handler

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
)

type Handler struct {
	runner  *runner.Runner
	store 	map[string]*runner.RCmd
	config  *config.Config
}

func New(r *runner.Runner, c *config.Config) *Handler {
	return &Handler{r, make(map[string]*runner.RCmd), c}
}

func (h *Handler) Register(g *echo.Group) {
	h.registerRS(g)
}


func (h *Handler) Store(cmd *runner.RCmd) string {
	u := uuid.NewV4().String()
	h.store[u] = cmd
	return u
}

func (h *Handler) Get(uuid string) *runner.RCmd {
	return h.store[uuid]
}