BIN=pdoc

default: build

clean:
	go clean

build:
	go build -o $(BIN)
