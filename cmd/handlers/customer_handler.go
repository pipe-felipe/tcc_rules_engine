package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pipe-felipe/tcc_rules_engine/cmd/rules"

	"github.com/labstack/echo/v4"
)

func rulesHandler(dto *Customer) {
	dto.TransactionStatus = rules.ReproveByEmail(dto.Email)
}

func TransactionalDataHandler(echoContext echo.Context) error {
	c := Customer{}
	//
	//if err := echoContext.Bind(&c); err != nil {
	//	return err
	//}

	defer echoContext.Request().Body.Close()

	b, err := ioutil.ReadAll(echoContext.Request().Body)
	if err != nil {
		log.Println("Faile to reagin body", err)
		return echoContext.String(http.StatusInternalServerError, "DEU RUIM")
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		print("deu ruim no json")
	}

	//customer := CustomerDTO{
	//	Name:             c.Name,
	//	Email:            c.Email,
	//	Document:         c.Document,
	//	CreditCard:       c.CreditCard,
	//	Address:          c.Address,
	//	BirthDate:        c.BirthDate,
	//	Age:              c.Age,
	//	TransactionCount: c.TransactionCount,
	//	AllTransactions:  c.AllTransactions,
	//	TransactionValue: c.TransactionValue,
	//}
	rulesHandler(&c)
	returnToTccRandom(c)
	return echoContext.JSON(http.StatusOK, c)
}

func returnToTccRandom(dto Customer) {
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
		panic(err)
	}

	fmt.Println("Sent ro random: ", res["json"])
}
