.PHONY: lint
default:

.PHONY: lint
lint:
	golangci-lint run --fix .

lint-ci:
	golangci-lint run .

# make makemigration NAME=some_name
.PHONY: makemigration
makemigration:
	migrate create -ext sql -dir migrations $(NAME)

.PHONY: gen
gen-doc:
	swag i -g main.go --parseInternal --parseDependency

.PHONY: mock
mock:
	 mockery --all --output mocks

.PHONY: test
test:
	@go test --race --vet= ./... -v
