package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/service"
)

func CreateCustomer(e *echo.Echo) {
	e.POST("/engine/customer", service.GetCustomerFromTccRandom)
}
