PROJECT_NAME := "tcc_rules_engine"

.PHONY: make
make:
	go build -o ./build/$(PROJECT_NAME)

.PHONY: clean
clean:
	rm ./build/$(PROJECT_NAME)

.PHONY: run
run:
	./build/$(PROJECT_NAME)