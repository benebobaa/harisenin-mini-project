postgres:
	sudo docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:16-alpine

createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root harisenin_project

dropdb:
	sudo docker exec -it postgres16 dropdb harisenin_project

migratecreate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin_project?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin_project?sslmode=disable" -verbose down 1

migratefix:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/harisenin_project?sslmode=disable" -verbose force $(version)

test:
	go test -v -cover ./...
