MIGRATIONS_PATH="db/migrations"
UPL_DATABASE_DRIVER ?= postgres
UPL_DATABASE_HOST ?= localhost
UPL_DATABASE_PORT ?= 5432
UPL_DATABASE_USER ?= docker
UPL_DATABASE_PASSWORD ?= docker
UPL_DATABASE_SSL ?= disable
UPL_DATABASE_DATABASE = upl
DATABASE_DSN := "${UPL_DATABASE_DRIVER}://${UPL_DATABASE_USER}:${UPL_DATABASE_PASSWORD}@${UPL_DATABASE_HOST}:${UPL_DATABASE_PORT}/${UPL_DATABASE_DATABASE}?sslmode=${UPL_DATABASE_SSL}"

#############
# DATABASE  #
#############
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)

#############
# SERVER    #
#############
.PHONY: run
run:
	go run ./cmd/server/main.go