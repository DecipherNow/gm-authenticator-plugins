.PHONY: test
test:
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v ./...
