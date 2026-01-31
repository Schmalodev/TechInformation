build:
	go build ./src/cmd/ram-monitor

deploy:
	build
	./ram-monitor

unit-test:
	go test ./...