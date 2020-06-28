package main

import "pantrycontrol-backend/internal/application/routes"

func main() {

	router := routes.Route()
	router.Start(":8080")
}
