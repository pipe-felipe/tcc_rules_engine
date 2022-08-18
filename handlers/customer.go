package handlers

type Customer struct {
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Document          string     `json:"document"`
	CreditCard        CreditCard `json:"creditCard"`
	Address           Address    `json:"address"`
	BirthDate         string     `json:"birthDate"`
	TransactionValue  float64    `json:"transactionValue"`
	TransactionStatus string     `json:"transactionStatus"`
}
