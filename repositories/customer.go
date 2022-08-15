package repositories

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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

func (c *Customer) GetTransactionalData(echoContext echo.Context) (error, io.Reader) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			echoContext.Logger().Error(err)
		}
	}(echoContext.Request().Body)

	body, err := ioutil.ReadAll(echoContext.Request().Body)
	if err != nil {
		err := echoContext.JSON(http.StatusBadRequest, err)
		log.Printf("Failed to read body: %s", err)
		return echoContext.String(http.StatusInternalServerError, "Failed to read body"), nil
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("Failed to unmarshal body: %s", err)
		return echoContext.String(http.StatusInternalServerError, "Failed to unmarshal body"), nil
	}
	log.Printf("Customer received: %+v", c)

	return echoContext.JSON(http.StatusOK, c), bytes.NewReader(body)
}
