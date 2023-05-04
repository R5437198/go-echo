#!make
-include .env

migrate-create-%:
	migrate create -ext sql -dir db/migrations -seq $*

migrate-up:
	migrate -database ${DB_URL} -path ./db/migrations/ up

migrate-down:
	migrate -database ${DB_URL} -path ./db/migrations/ down -all

migrate-force-%:
	migrate -database ${DB_URL} -path ./db/migrations/ force $*

init:
	make migrate-down
	make migrate-up
