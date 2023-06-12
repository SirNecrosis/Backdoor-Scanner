.PHONY: run
run:
	go run main.go

.PHONY: bin
bin:
	GOOS=linux GOARCH=amd64 go build -o bin/BackdoorScanner main.go
	GOOS=windows GOARCH=amd64 go build -o bin/BackdoorScanner.exe main.go

.PHONY: tidy
tidy:
	go mod tidy