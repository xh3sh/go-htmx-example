package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xh3sh/go-htmx-example/internal/config"
	"github.com/xh3sh/go-htmx-example/internal/handlers"
)

func main() {

	e := echo.New()
	e.Renderer = handlers.NewTemplates()
	e.Use(middleware.Logger())

	// Load config
	config, err := config.LoadConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	fs := http.FileServer(http.Dir("static"))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", fs)))

	// Pass config to handler
	h := handlers.NewHandler(config)

	e.GET("/", h.HandleHome)
	e.GET("/filter", h.HandleFilter)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
