MIGRATIONS_PATH="db/migrations"

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

.PHONY: run
run:
	go run ./cmd/server/main.go