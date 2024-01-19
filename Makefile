migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin?sslmode=disable" -verbose down 1

migratefix:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin?sslmode=disable" -verbose force $(version)

test:
	go test -v -cover ./...

PHONY: postgres createdb dropdb migrate_create migrate_up migrate_down test migrate_fix