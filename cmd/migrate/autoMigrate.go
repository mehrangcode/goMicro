package main

import (
	"mehrangcode.ir/todoapp/cmd/initializers"
	"mehrangcode.ir/todoapp/cmd/models"
)

func init() {
	// initializers.LoadEnvFile()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
