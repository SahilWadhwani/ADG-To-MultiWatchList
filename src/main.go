package main

import (
	"main/src/database"
	"main/src/routes"
)

func main() {

	database.InitDB()

	r := routes.SetupRouter()
	r.Run(":8080")

}
