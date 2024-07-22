package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spilikin/go-monorepo/lib"
)

func main() {
	fmt.Println(lib.MessageFromLib())
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
