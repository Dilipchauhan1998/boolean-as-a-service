package main

import (
	"boolean-as-a-service/conn"
	"boolean-as-a-service/models"
	"boolean-as-a-service/routes"
)

func main() {
	conn.ConnectDB()
	db := conn.DB
	db.AutoMigrate(&models.Boolean{})
	defer db.Close()

	router := routes.SetupRouter()
	router.Run(":80")
}
