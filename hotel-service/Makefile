run:
	go run cmd/soldiers_service/main.go



migrate-file:
	migrate create -ext sql -dir migrations/ -seq room

DB_URL := "postgres://postgres:+_+diyor2005+_+@hotel_postgres:5432/hotel?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1