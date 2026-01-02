DB_URL=postgres://postgres:postgres@localhost:5432/shortener?sslmode=disable

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down 1

migrate-status:
	migrate -path migrations -database $(DB_URL) version
