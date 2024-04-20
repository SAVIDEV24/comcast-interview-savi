package main

import (
	"fmt"
	"net/http"
	"stringinator-go/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

	var stringmanipulator = service.StringinatorService{Seen_strings: make(map[string]int)}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<pre>
			Welcome to the Stringinator 3000 for all of your string manipulation needs.
			GET / - You're already here!
			POST /stringinate - Get all of the info you've ever wanted about a string. Takes JSON of the following form: {"input":"your-string-goes-here"}
			GET /stats - Get statistics about all strings the server has seen, including the longest and most popular strings.
			</pre>
		`)
	})

	e.POST("/stringinate", stringmanipulator.Stringinate)
	e.GET("/stringinate", stringmanipulator.Stringinate)
	e.GET("/stats", stringmanipulator.Stats)
	e.Logger.Fatal(e.Start(":1323"))
}
