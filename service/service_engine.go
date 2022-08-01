package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pipe-felipe/tcc_rules_engine/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Isso é o que eu vou fazer para o Random POST
// Essa função vai ser chamada quando todo o dado for processado pelas regras
func SendCurateDataToRandom(url string) {
	requestBody := strings.NewReader(`
		{
		  "name": "tcc_rule_engine2",
		  "document": "0101010101010100",
		  "email": "tcc_rule_engine2@tcc.com",
		  "creditCard": {
			"flag": "VISA",
			"holderName": "Felipe A Nascimento",
			"number": "3215123565487899",
			"validThru": "2021-10",
			"cvv": "215"
		  },
		  "address": {
			"street": "Go land",
			"number": 10101001,
			"city": "Google",
			"state": "São Paulo",
			"country": "Brazil",
			"zipCode": "13253-710"
		  },
		  "birthDate": "2009-12-20",
		  "transactionValue": 123.45
		}
	`)
	response, err := http.Post(url, "application/json", requestBody)

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
}

func GetCustomerFromTccRandom(echoContext echo.Context) error {
	customer := model.Customer{}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			echoContext.Logger().Error(err)
		}
	}(echoContext.Request().Body)

	body, err := ioutil.ReadAll(echoContext.Request().Body)
	if err != nil {
		err := echoContext.JSON(http.StatusBadRequest, err)
		if err != nil {
			return nil
		}
		log.Printf("Failed to read body: %s", err)
		return echoContext.String(http.StatusInternalServerError, "Failed to read body")
	}

	err = json.Unmarshal(body, &customer)
	if err != nil {
		log.Printf("Failed to unmarshal body: %s", err)
		return echoContext.String(http.StatusInternalServerError, "Failed to unmarshal body")
	}

	log.Printf("Customer: %+v", customer)
	return echoContext.JSON(http.StatusOK, customer)
}
