run-local:
	go run ./cmd/pantry-control/main.go

run-migration:
	go run ./tools/database/migration.go