package handlers

type CustomerDTO struct {
	Name              string
	Email             string
	Document          string
	CreditCard        CreditCard
	Address           Address
	BirthDate         string
	TransactionValue  float64
	TransactionStatus string
}
