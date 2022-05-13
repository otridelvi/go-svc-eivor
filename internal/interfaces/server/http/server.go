package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/otridelvi/go-svc-eivor/internal/interfaces/server/container"
	v1 "github.com/otridelvi/go-svc-eivor/internal/interfaces/server/http/api/v1"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartHttpService(cont *container.Container) {
	server := echo.New()
	server.HideBanner = true

	v1Router := v1.NewRouter(server, cont)
	v1Router.RegisterRoutes()

	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})

	// start server
	go func() {
		if err := server.Start(cont.Config.AppAddress()); err != nil {
			server.Logger.Print(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}

}
