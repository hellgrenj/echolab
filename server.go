package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/hellgrenj/echolab/features/about"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"html/template"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	if string(name[0]) == "_" { //_partialsStartsWithUnderScore ... not based on base.html
		return tmpl.ExecuteTemplate(w, name, data)
	} else {
		m, ok := data.(map[string]interface{})
		if ok == true {
			m["isProd"] = false // ..base on env var.. and use minify as part of docker build to se prod and bundle js and css...
		}
		return tmpl.ExecuteTemplate(w, "base.html", data) // base-based-page :)
	}

}

func main() {

	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("features/home/home.html", "features/shared/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("features/about/about.html", "features/shared/base.html"))
	templates["_somePartial.html"] = template.Must(template.ParseFiles("features/shared/_somePartial.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.File("features/shared/util.js", "features/shared/util.js")
	e.GET("/", Home)
	e.GET("/home", Home)
	about.Register(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Home",
		"msg":  "the home page",
	})
}
