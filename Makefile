build:
	go build ./src/cmd/ram-monitor

run:
	go run ./cmd/ram-monitor

unit-test:
	go test ./...