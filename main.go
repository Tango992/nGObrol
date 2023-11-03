package main

import (
	"avengers-chat/controller"
	"avengers-chat/handler"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Pre(middleware.HTTPSRedirect())


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

	e.GET("/ws/inventory", controller.InventoryWebsocket)
	e.GET("/ws/crime", controller.CrimeWebsocket)
	e.GET("/ws/random", controller.RandomWebsocket)

	// e.Logger.Fatal(e.Start(os.Getenv("PORT")))
	if err := e.StartTLS(os.Getenv("PORT"), []byte(os.Getenv("TSL_CERT")), []byte(os.Getenv("TSL_KEY"))); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
