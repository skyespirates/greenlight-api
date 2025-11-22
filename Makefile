migrations_path = ./migrations
dsn = postgres://postgres:bmwb1gtr@localhost:5432/clean?sslmode=disable

.PHONY: api/start
api/start:
	@go run .\cmd\api

.PHONY: api/dev
api/dev:
	@air

.PHONY: db/migrations/new
db/migrations/new:
	@migrate create -seq -ext sql -dir $(migrations_path) $(name)

.PHONY: db/migrations/new
db/migrations/up:
	@migrate -database $(dsn) -path $(migrations_path) up

.PHONY: db/migrations/down
db/migrations/down:
	@migrate -database $(dsn) -path $(migrations_path) down

.PHONY: db/migrations/version
db/migrations/version:
	@migrate -database $(dsn) -path $(migrations_path) version