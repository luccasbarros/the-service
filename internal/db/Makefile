.PHONY: migrate-create migrate-up migrate-down

migrate-create:
	migrate create -ext sql -dir internal/db/migrations/ -seq $(name)

migrate-up:
	migrate -database "postgres://master:master@localhost:5433/service?sslmode=disable" -path migrations up

migrate-down:
	migrate -database "postgres://master:master@localhost:5433/service?sslmode=disable" -path migrations down 1