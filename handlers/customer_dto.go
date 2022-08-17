package handlers

type CustomerDTO struct {
	Name              string
	Document          string
	Email             string
	CreditCard        CreditCard
	Address           Address
	BirthDate         string
	TransactionValue  float64
	TransactionStatus string
}
