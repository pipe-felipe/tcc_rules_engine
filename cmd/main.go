package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/cmd/controller"
)

func main() {
	e := echo.New()
	controller.TransactionDataRetriever(e)
	err := e.Start(":8082")
	if err != nil {
		panic(err)
	}
}
