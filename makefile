MAIN = "cmd/main.go"
PGDNS = "host=localhost user=postgres password=postgres dbname=auth port=5432 sslmode=disable"

generate:
	protoc --go_out=. --go-grpc_out=. api/api.proto

mockgen:
	mockgen -destination=internal/pkg/mock/store.go -package=mock auth-service/internal/pkg/store IStore
	mockgen -destination=internal/pkg/mock/jwt.go -package=mock auth-service/internal/pkg/encrypt IJWT

run:
	go run $(MAIN)

db-start:
	docker start /auth-service-postgres

db-up:
	goose -dir ./migrations postgres $(PGDNS) up

db-down:
	goose -dir ./migrations postgres $(PGDNS) down

db-create:
	goose -dir ./migrations postgres $(PGDNS) create _ sql

dc-up:
	docker run -d --name auth-service-postgres -p 5432:5432 -e POSTGRES_DB=auth -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:15-alpine
