all: test

test: build
	@./cmdapp -l
	@./cmdapp -n D01

build:
	@go build

clean:
	@go clean