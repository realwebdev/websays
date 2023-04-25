postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres clockify

dropdb:
	docker exec -it postgres12 dropdb clockify

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb test
