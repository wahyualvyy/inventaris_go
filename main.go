package main

import (
	"lab-inventaris/config"
	"lab-inventaris/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()

	r.Run(":8080")
}