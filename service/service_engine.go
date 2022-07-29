package service

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

		}
	}(response.Body)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", content)
}
