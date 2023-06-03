.PHONY: run
run:
	go run BackdoorScanner.go

.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 go build -o bin/BackdoorScanner BackdoorScanner.go