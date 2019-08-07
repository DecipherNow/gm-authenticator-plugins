.PHONY: test
test:
	@echo "--> Running tests..."
	@CGO_ENABLED=0 go test -v ./...

.PHONY: vendor
vendor:
	@echo "--> Vendoring dependencies..."
	@CGO_ENABLED=0 go mod vendor
