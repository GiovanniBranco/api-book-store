package main

import (
	"github.com/GiovanniBranco/classroom-api/database"
	"github.com/GiovanniBranco/classroom-api/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequests()
}
