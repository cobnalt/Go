PROGECT_PATH_WIN=C:\GoTest
PROGECT_PATH_LINUX=$(shell pwd) 

postgres:
	docker run --name postgres1 -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres1 createdb --username=postgres --owner=postgres simple_catalog

dropdb:
	docker exec -it postgres1 dropdb -U postgres simple_catalog

migrateup:	
	docker run -v $(PROGECT_PATH_WIN)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://postgres:postgres@localhost:5432/simple_catalog?sslmode=disable up

migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/simple_catalog?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown 