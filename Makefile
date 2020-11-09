postgres:
	docker run --name postgres1 -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres1 createdb --username=postgres --owner=postgres simple_catalog

dropdb:
	docker exec -it postgres1 dropdb -U postgres simple_catalog

migrateup:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/simple_catalog?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/simple_catalog?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown 