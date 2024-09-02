all: mock lint test build

mock:
	mockery --dir ./internal/dao --output ./internal/dao/mocks --all --disable-version-string
	mockery --dir ./internal/account --output ./internal/account/mocks --all --disable-version-string

lint:
	go mod tidy
	golangci-lint run

build:
	go mod tidy
	go build ./cmd/main.go

run:
	go mod tidy
	env=local go run ./cmd/main.go

test:
	go mod tidy
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out