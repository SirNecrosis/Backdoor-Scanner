.PHONY: run
run:
	go run cmd/BackdoorScanner.go

.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 go build -o bin/BackdoorScanner cmd/BackdoorScanner.go
	GOOS=windows GOARCH=amd64 go build -o bin/BackdoorScanner.exe cmd/BackdoorScanner.go

.PHONY: tidy
tidy:
	go mod tidy