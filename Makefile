.PHONY:

COV_CUTOVER = 0

## Add test user
devtools-add-test-user:
	go run devtools/add-user.go

## Bring up dev server
dev-backend-run:
	go run ./backend

## Bring up docker database
dev-db-up:
	docker-compose -f ./docker/docker-compose.yml --env-file .env up 

install: 
	go get -u ./...

## test: run unit/mock tests
test: generate
	go test v -cover ./...

## test-only: run unit/mock tests without any dependent step
test-only:
	go test -v -cover ./...

## test-race: run go unit/mock tests with race detection
test-race: generate
	go test -v -race ./...

## test-race-only: run go unit/mock tests with race detection without any dependent step
test-race-only:
	go test -v -race ./...

## test-integration: runs go integration tests
test-integration: generate
	go test -v -tags=integration ./...

## test-integration-only: runs go integration tests without any dependent step
test-integration-only:
	go test -v -tags=integration ./...

## test-all: runs all tests
test-all: generate
	go test -v ./...
	go test -v -tags=integration ./...

## cover: run unit/mock tests with coverage report. Generated mocks are filtered out of the report
cover: generate
	go test -v --race -coverprofile=coverage.out -coverpkg=./... ./...
	cat coverage.out | grep -v "mock" > coverage.nomocks.out
	go tool cover -func coverage.nomocks.out

## cover-check: checks the code coverage to be beyond a certain threshold
cover-check: cover
	COV=$$(go tool cover -func coverage.nomocks.out | awk 'END{print substr($$3, 1, length($$3)-1)}') ;\
	echo "Test coverage at $$COV %" ;\
	if [ $$(echo "$$COV < $(COV_CUTOVER)" | bc) -eq 1 ]; then echo "Failing due to low test coverage"; exit 1; fi\

## lint: runs golangci-lint
lint: generate
	golangci-lint run -v ./...

## lint-only: runs golangci-lint without any dependent step
lint-only:
	golangci-lint run -v ./...
## generate: runs go generate
generate:
	go generate -v ./...

