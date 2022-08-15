package main

import (
	//tcc "github.com/pipe-felipe/tcc_rules_engine/tcc_random_handler"
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/controller"
	"github.com/pipe-felipe/tcc_rules_engine/repositories"
)

func main() {
	e := echo.New()
	c := repositories.Customer{}
	controller.CreateCustomer(e, &c)
	err := e.Start(":8082")
	if err != nil {
		return
	}
}
