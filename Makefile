tidy:
	@go mod tidy 

run:
	@go run cmd/main.go 

migrate:	
	migrate create -dir ./db -ext sql db

migrate-up:
	migrate -path ./db -database "postgres://postgres:14022014@localhost:5432/fitness_traking?sslmode=disable" up

migrate-down:
	migrate -path ./db -database "postgres://postgres:14022014@localhost:5432/fitness_traking?sslmode=disable" down

migrate-force:
	migrate -path ./db -database "postgres://postgres:14022014@localhost:5432/fitness_traking?sslmode=disable" force
