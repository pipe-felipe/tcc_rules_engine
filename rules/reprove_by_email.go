package rules

import (
	"github.com/pipe-felipe/tcc_rules_engine/model"
)

func ReproveByEmail(customer model.Customer) {
	if customer.Email == "salafrario@ladrão.com" {
		customer.TransactionStatus = "reproved"
	}
}
