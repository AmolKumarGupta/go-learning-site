BINARY_NAME=app

build: 
	go build -o bin/${BINARY_NAME} main.go

run: build
	./bin/${BINARY_NAME}