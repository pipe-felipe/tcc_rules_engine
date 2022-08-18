package rules

func ReproveByEmail(email string) string {
	if email != "salafrario@ladrão.com" {
		return "APROVED"
	} else {
		return "REPROVED"
	}
}
