package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/rules"
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
		Name:             c.Name,
		Email:            c.Email,
		Document:         c.Document,
		CreditCard:       c.CreditCard,
		Address:          c.Address,
		BirthDate:        c.BirthDate,
		TransactionValue: c.TransactionValue,
	}
	rulesHandler(&customer)
	SendTransactionalData(&customer)
	return echoContext.JSON(http.StatusOK, customer)
}

func SendTransactionalData(dto *CustomerDTO) {
	URL := "http://localhost:8080/customer"

	jsonData, err := json.Marshal(dto)
	responseBody := bytes.NewBuffer(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.NewRequest("POST", URL, responseBody)
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	post := &CustomerDTO{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}
	log.Println("Final: ", post)
}
