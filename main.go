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
	models.SetBooleanRepo(conn.DB)

	router := routes.SetupRouter()
	router.Run(":80")
}
