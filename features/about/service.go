package about

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageModel struct {
	Title string
}

func Register(e *echo.Echo) {
	e.GET("/about", PageHandler)
	e.GET("/about/api/somesome", SomeAsyncStuff)
	e.GET("/about/somep", SomePartial)
	e.File("/about/about.js", "features/about/about.js")
}
func PageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"msg":  "the about page",
	})
}
func SomeAsyncStuff(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	})
}
func SomePartial(c echo.Context) error {
	return c.Render(http.StatusOK, "_somePartial.html", map[string]interface{}{
		"msg": "this is a partial",
	})
}
