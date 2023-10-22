package main

import (
	"html/template"
	"io"
	"log"

	"manu/htmx.conway/pkg/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	tmpl, err := template.ParseGlob(
		"../public/views/*.html",
	)
	if err != nil {
		log.Fatal("Could not load template files: ", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}

	e.Use(middleware.Logger())
	e.Static("/css", "../css")

	e.GET("/", pages.Index)
	e.POST("/kill/:PosX/:PosY", pages.Kill)
	e.POST("/alive/:PosX/:PosY", pages.Alive)
	e.POST("/advance", pages.Advance)
	e.POST("/reset", pages.Reset)
	e.PUT("/start", pages.Start)

	e.Logger.Fatal(e.Start(":42069"))
}
