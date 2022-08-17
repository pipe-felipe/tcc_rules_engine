package handlers

import (
	"github.com/pipe-felipe/tcc_rules_engine/rules"
)

//func (c *Customer) CustomerHandler(e echo.Context) error {
//	url := "http://localhost:8080/customer"
//	_, body := c.GetTransactionalData(e)
//	c.rulesHandler()
//	response, err := http.Post(url, "application/json", body)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(response.Body)
//
//	content, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		panic(err)
//	}
//
//	log.Printf("%s", content)
//	return e.JSON(http.StatusOK, body)
//}

func (c *Customer) rulesHandler() {
	c.TransactionStatus = rules.ReproveByEmail(c.Email)
}
