all:

test:
	@make -C demos/cmdapp $@

clean:
	@make -C demos/cmdapp $@
	@rm -f output.*

cov:
	@go test -coverprofile=output.cov
	@go tool cover -func=output.cov

doc:
	@go tool doc -all
	@go tool doc -C demos/cmdapp -cmd -all