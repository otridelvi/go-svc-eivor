package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/otridelvi/go-svc-eivor/internal/interfaces/server/container"
	"net/http"
)

type Router struct {
	V1Group *echo.Group
}

func NewRouter(server *echo.Echo, cont *container.Container) *Router {
	return &Router{
		V1Group: server.Group("v1"),
	}
}

func (r *Router) RegisterRoutes() {
	r.V1Group.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Service is up")
	})
}
