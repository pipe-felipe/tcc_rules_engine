package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/handlers"
)

func TransactionDataRetriever(e *echo.Echo) {
	e.POST("/engine/customer", handlers.GetTransactionalData)
}
