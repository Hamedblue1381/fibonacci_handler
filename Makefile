server:
	@go build -o bin/fibonacci main.go
	@./bin/fibonacci
test:
	@go test ./... -v
