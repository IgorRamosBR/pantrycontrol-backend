package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"pantrycontrol-backend/internal/infra/configuration"
)

func main() {
	appConfig := configuration.CreateConfig()

	m, err := migrate.New(
		"file://db/migrations",
		appConfig.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
