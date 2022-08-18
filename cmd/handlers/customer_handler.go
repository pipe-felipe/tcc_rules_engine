package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pipe-felipe/tcc_rules_engine/cmd/rules"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	returnToTccRandom(customer)
	return echoContext.JSON(http.StatusOK, customer)
}

func returnToTccRandom(dto CustomerDTO) {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		log.Fatal("Error on marshal: ", err)
	}

	resp, err := http.Post(TccRandomUrl, "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal("Error on POST the data to tcc_random", err)
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}

	fmt.Println("Sent ro random: ", res["json"])
}
