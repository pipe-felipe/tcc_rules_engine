package handlers

type CustomerDTO struct {
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Document          string     `json:"document"`
	CreditCard        CreditCard `json:"creditCard"`
	Address           Address    `json:"address"`
	BirthDate         string     `json:"birthDate"`
	Age               int        `json:"age"`
	TransactionValue  float64    `json:"transactionValue"`
	TransactionCount  int        `json:"transactionCount"`
	AllTransactions   []float64  `json:"allTransactions"`
	TransactionStatus string     `json:"transactionStatus"`
}
