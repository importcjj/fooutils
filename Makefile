help:
	@echo "Commands:"
	@echo "  test:  Run test case."

test:
	go test -v ./...

install:
	go get ./...
