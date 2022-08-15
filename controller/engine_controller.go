package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/repositories"
)

func CreateCustomer(e *echo.Echo, c *repositories.Customer) {
	e.POST("/engine/customer", c.CustomerHandler)
}
