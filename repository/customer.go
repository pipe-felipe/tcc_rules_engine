package repository

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"log"
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

func (c *Customer) ReadTest(e echo.Context) error {
	url := "http://localhost:8080/customer"
	err, body := c.GetTransactionalData(e)
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", content)
	return e.JSON(http.StatusOK, body)
}
