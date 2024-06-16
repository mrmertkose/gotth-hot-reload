package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrmertkose/gotth-hot-reload/views"
)

var dev = true

func main() {
	e := echo.New()
	e.Static("/assets", "./assets")
	e.GET("/", homeHandler)
	log.Fatal(e.Start(":3000"))
}

func homeHandler(c echo.Context) error {
	return render(c, views.Index())
}

func render(c echo.Context, component templ.Component) error {
	if dev {
		c.Response().Header().Set("Cache-Control", "no-cache, private, max-age=0")
	}
	return component.Render(c.Request().Context(), c.Response())
}
