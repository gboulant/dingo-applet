all: test

demos.%:
	@make -C demos/d01.basic $*
	@make -C demos/d02.applet $*

test: demos.test

clean: demos.clean
	@rm -f output.*

cov:
	@go test -coverprofile=output.cov
	@go tool cover -func=output.cov

doc:
	@go tool doc -all
	@go tool doc -C demos/cmdapp -cmd -all