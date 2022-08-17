package handlers

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
