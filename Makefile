test:
	go test $$(go list ./... | grep -v /vendor/)

run-local:
	go run ./cmd/pantry-control/main.go

run-migration:
	go run ./tools/database/migration.go