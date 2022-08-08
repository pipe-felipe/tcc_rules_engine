package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/repository"
)

func CreateCustomer(e *echo.Echo, c *repository.Customer) {
	e.POST("/engine/customer", c.ReadTest)
}
