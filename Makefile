build:
	@go build -o bin/Library src/cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/Library

migration:
	@migrate create -ext sql -dir src/cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run src/cmd/migrate/main.go up

migrate-down:
	@go run src/cmd/migrate/main.go down
