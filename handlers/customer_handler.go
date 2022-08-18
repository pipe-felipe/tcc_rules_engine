package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		Name: c.Name,
	}
	rulesHandler(&customer)
	SendTransactionalData(customer)
	return echoContext.JSON(http.StatusOK, customer)
}

func SendTransactionalData(dto CustomerDTO) {
	URL := "http://localhost:8080/customer/" + dto.Document

	json_data, err := json.Marshal(dto)
	if err != nil {
		log.Fatal("Error on unmarshal: ", err)
	}

	resp, err := http.Post(URL, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal("Error on POST the data to tcc_random", err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println("Sent ro random: ", res["json"])
}
