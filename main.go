package main

import (
	//tcc "github.com/pipe-felipe/tcc_rules_engine/tcc_random_handler"
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/controller"
)

func main() {
	e := echo.New()
	controller.CreateCustomer(e)

	err := e.Start(":8082")
	if err != nil {
		return
	}
}
