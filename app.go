package main

import (
	"net/http"
	"stringinator-go/datastore"
	"stringinator-go/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//persistStore := datastore.NewInMemoryStore(model.FilePath)
	tempStore := datastore.NewTempIms(make(map[string]int))
	var stringmanipulator = service.NewStringinatorService(tempStore)

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
