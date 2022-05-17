container: 
	docker run --name pokemonAPI -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pokemonAPI createdb --username=root --owner=root pokemon

dropdb:
	docker exec -it pokemonAPI dropdb --username=root pokemon

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/pokemon?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/pokemon?sslmode=disable" -verbose down
	
sqlc:
	docker run --rm -v C:\Users\amir\.vscode\interviews\EcoChain\pokemon-api-hjfxii:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb container dropdb migrateup migratedown sqlc test server