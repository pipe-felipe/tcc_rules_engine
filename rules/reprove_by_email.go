package rules

import (
	"github.com/pipe-felipe/tcc_rules_engine/repository"
)

func ReproveByEmail(customer *repository.Customer) {
	if customer.Email == "salafrario@ladrão.com" {
		customer.TransactionStatus = "REPROVE"
	} else {
		customer.TransactionStatus = "APPROVE"
	}
}
