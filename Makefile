.PHONY: create

migrate-create:
	migrate create -ext sql -dir internal/db/migrations/ -seq $(name)

migrate-up:
	migrate -database "postgres://master:master@localhost:5433/service?sslmode=disable" -path internal/db/migrations/ up

migrate-down:
	migrate -database "postgres://master:master@localhost:5433/service?sslmode=disable" -path internal/db/migrations down 1