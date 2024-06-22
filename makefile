include .env

run:
	docker-compose up -d

dev:
	docker-compose up -d && make logs

logs:
	docker logs --follow docker_cinema_api

migration:
	docker exec -it docker_cinema_api migrate create -ext sql -dir internal/database/migrations -seq $(name)

migrate:
	docker exec -it docker_cinema_api migrate -path app/internal/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-undo:
	docker exec -it docker_cinema_api migrate -path app/internal/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down -all
