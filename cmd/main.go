package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pipe-felipe/tcc_rules_engine/cmd/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	controller.TransactionDataRetriever(e)
	err := e.Start(":8082")

	if err != nil {
		panic(err)
	}
}
