MAIN = "cmd/main.go"
PGDNS = "host=localhost user=postgres password=postgres dbname=auth port=5432 sslmode=disable"

generate:
	protoc --go_out=. --go-grpc_out=. api/api.proto

make run:
	go run $(MAIN)

make db-start:
	docker start /auth-service-postgres

make db-up:
	goose -dir ./migrations postgres $(PGDNS) up

make db-down:
	goose -dir ./migrations postgres $(PGDNS) down

make db-create:
	goose -dir ./migrations postgres $(PGDNS) create _ sql

make dc-up:
	docker run -d --name auth-service-postgres -p 5432:5432 -e POSTGRES_DB=auth -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:15-alpine
