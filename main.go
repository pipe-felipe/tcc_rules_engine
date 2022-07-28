package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func handlerTccRandom(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello from nothing"})
}

func main() {
	e := echo.New()
	e.GET("/", handlerTccRandom)
	e.Logger.Fatal(e.Start(":1323"))
}
