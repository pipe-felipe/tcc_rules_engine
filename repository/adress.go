package repository

type Address struct {
	Street  string `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}
