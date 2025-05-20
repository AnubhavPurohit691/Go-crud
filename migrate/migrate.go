package main

import (
	"github.com/AnubhavPurohit691/go-crud/initializer"
	"github.com/AnubhavPurohit691/go-crud/models"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Todo{})
}
