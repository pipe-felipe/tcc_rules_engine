package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/rules"
	"net/http"
)

func rulesHandler(dto *CustomerDTO) {
	dto.TransactionStatus = rules.ReproveByEmail(dto.Email)
}

func TransactionalDataHandler(echoContext echo.Context) error {
	c := new(Customer)
	if err := echoContext.Bind(c); err != nil {
		return nil
	}
	customer := CustomerDTO{
		Name: c.Name,
	}
	rulesHandler(&customer)
	return echoContext.JSON(http.StatusOK, customer)
}
