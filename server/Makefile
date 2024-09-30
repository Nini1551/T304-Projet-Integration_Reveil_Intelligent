postgresinit:
	docker compose up -d postgres

postgresdown:
	docker compose down

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root alarm-clock

dropdb:
	docker exec -it postgres dropdb alarm-clock

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/alarm-clock?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/alarm-clock?sslmode=disable" -verbose down

.PHONY: postgresinit postgresdown postgres createdb dropdb migrateup migratedown