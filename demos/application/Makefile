all: test

test: build
	@./application -l
	@./application -n D01

build:
	@go build

clean:
	@go clean