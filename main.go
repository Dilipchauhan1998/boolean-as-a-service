package main

import (
	"boolean-as-a-service/auth"
	"boolean-as-a-service/conn"
	"boolean-as-a-service/models"
	"boolean-as-a-service/routes"
)

func main() {
	conn.ConnectDB()
	db := conn.DB
	db.AutoMigrate(&models.Boolean{})
	db.AutoMigrate(&auth.Token{})
	defer db.Close()
	models.SetBooleanRepo(conn.DB)
	auth.SetTokenRepo(conn.DB)

	router := routes.SetupRouter()
	router.Run(":80")
}
