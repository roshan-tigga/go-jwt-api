package main

import (
	"auth/config"
	"auth/routes"
	"fmt"
)

func main() {
	config.DataMigration()
	router := routes.InitializeRoutes()
	port := 8000
	fmt.Printf("Server started on :%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}
