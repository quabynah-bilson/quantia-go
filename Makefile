# set the BINARY_NAME variable to the name of our binary
BINARY_NAME = server

clean:
	@rm -f ./bin/$(BINARY_NAME)

build: clean
	@go build -o ./bin/$(BINARY_NAME) -v ./cmd/...

run-tests: build
	@go test -cover -v ./tests/...

run: build
	@./bin/$(BINARY_NAME)

run-databases:
	@docker-compose up -d --remove-orphans

stop-databases:
	@docker-compose down

coverage:
	@go test -coverprofile=coverage/coverage.txt -covermode count ./tests/...
	@gocover-cobertura < coverage/coverage.txt > coverage/coverage.cobertura.xml

.PHONY: clean build run-tests run run-databases stop-databases run-lints coverage