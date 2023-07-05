.PHONY: create

create:
	migrate create -ext sql -dir internal/db/migrations/ -seq $(name)