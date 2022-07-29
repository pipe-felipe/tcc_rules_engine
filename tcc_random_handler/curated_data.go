package tcc_random_handler

import (
	"tcc_rules_engine/service"
)

func EngineReturn() {
	service.SendCurateDataToRandom("http://localhost:8080/customer")
}
