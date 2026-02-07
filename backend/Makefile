.PHONY: run build swagger

run:
	go run ./cmd/server/main.go

build:
	go build ./cmd/server/main.go 

swagger:
	swag init -g cmd/server/main.go -o docs

init:
	go install github.com/swaggo/swag/cmd/swag@latest 
	go mod download
	go mod tidy