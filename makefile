output=./bin/ricky

.PHONY: build run

build:
	go build -o ${output}
	${output}

run: 
	go run .
