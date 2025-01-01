migrate-add:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-run:
	migrate -database postgres://greenlight:greenlight@0.0.0.0:5432/greenlight?sslmode=disable -path migrations up