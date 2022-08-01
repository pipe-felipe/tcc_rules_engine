package model

type Customer struct {
	Name       string `json:"name"`
	Document   string `json:"document"`
	Email      string `json:"email"`
	CreditCard struct {
		Flag      string `json:"flag"`
		Holder    string `json:"holderName"`
		Number    string `json:"number"`
		ValidThru string `json:"validThru"`
		CVV       string `json:"cvv"`
	}
	Address struct {
		Street  string `json:"street"`
		Number  int    `json:"number"`
		City    string `json:"city"`
		State   string `json:"state"`
		Country string `json:"country"`
		ZipCode string `json:"zipCode"`
	}
	BirthDate        string  `json:"birthDate"`
	TransactionValue float64 `json:"transactionValue"`
}
