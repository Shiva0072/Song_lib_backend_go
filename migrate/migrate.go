package main

import (
	"learnGo/Crud_CC/entities"
	"learnGo/Crud_CC/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&entities.Song{})
}
