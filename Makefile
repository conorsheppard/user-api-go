SHELL:=/bin/bash

DB_URL=postgresql://root:secret@localhost:5432/user_api_go?sslmode=disable

default: build run

network:
	docker network create user-network

build:
	docker build -t user-api-go .

run: build
	docker run --rm --network user-network -p 8080:8080 -t user-api-go

local-run:
	go run main.go

postgres:
	docker run --name postgres --network user-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

create-db:
	docker exec -it postgres createdb --username=root --owner=root user_api_go

migrate-up:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

sqlc: # looks at internal/db/query/user.sql to generate source code
	sqlc generate

up:
	docker-compose up
	sleep 1
	make create-db
	make migrate-up

down:
	docker-compose down
	# docker kill $$(docker ps -aqf "name=postgres")
	# echo 'y' | docker container prune
	 docker image rm $$(docker images --format="{{.Repository}} {{.ID}}" | grep "^user-api-go_api " | cut -d' ' -f2)

ssh-db:
	docker exec -it $$(docker ps -aqf "name=postgres") psql -d user_api_go -U root

mockdb:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/conorsheppard/user-api-go/internal/db/sqlc Store

test:
	go test -v ./...

.PHONY: build run local-run postgres create-db migrate-up migrate-down sqlc compose-up compose-down ssh-db mock