package main

import (
	"devmoobin/go-crud/initializers"
	"devmoobin/go-crud/models"
)

func init() {
	
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
   initializers.DB.AutoMigrate(&models.Post{})
}
