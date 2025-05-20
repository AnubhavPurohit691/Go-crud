package main

import (
	"github.com/AnubhavPurohit691/go-crud/initializer"
	"github.com/AnubhavPurohit691/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()

}

func main() {
	r := gin.Default()

	routes.TodoRoutes(r)
	r.Run()
}
