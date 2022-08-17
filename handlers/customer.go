package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Customer struct {
	Name              string     `json:"name"`
	Document          string     `json:"document"`
	Email             string     `json:"email"`
	CreditCard        CreditCard `json:"creditCard"`
	Address           Address    `json:"address"`
	BirthDate         string     `json:"birthDate"`
	TransactionValue  float64    `json:"transactionValue"`
	TransactionStatus string     `json:"transactionStatus"`
}

type CustomerDTO struct {
	Name string
}

func GetTransactionalData(echoContext echo.Context) error {
	c := new(Customer)
	if err := echoContext.Bind(c); err != nil {
		return nil
	}
	customer := CustomerDTO{
		Name: c.Name,
	}
	//executeSomeBusinessLogic(customer)
	return echoContext.JSON(http.StatusOK, customer)
}
