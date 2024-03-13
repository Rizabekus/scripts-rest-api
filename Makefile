run:
	go run ./cmd/web
migrate:
	psql -U postgres -c "CREATE DATABASE script;"
	