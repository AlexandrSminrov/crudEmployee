include env
export $(shell sed 's/=.*//' env)

run: ## run main.go
	@go run cmd/main.go