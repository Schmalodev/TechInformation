build:
	go build ./src/cmd/ram-monitor

deploy:
	make build
	./ram-monitor

unit-test:
	go test ./...