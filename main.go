package main

import (
	"avengers-chat/controller"
	"avengers-chat/handler"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", handler.Menu)
	e.GET("/inventory", handler.InventoryChatRoom)
	e.GET("/crime", handler.CrimesChatRoom)
	e.GET("/random", handler.RandomChatRoom)

	e.GET("inventory/ws", controller.InventoryWebsocket)
	e.GET("crime/ws", controller.CrimeWebsocket)
	e.GET("random/ws", controller.RandomWebsocket)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
