package handlers

type CreditCard struct {
	Flag      string `json:"flag"`
	Holder    string `json:"holderName"`
	Number    string `json:"number"`
	ValidThru string `json:"validThru"`
	CVV       string `json:"cvv"`
}
