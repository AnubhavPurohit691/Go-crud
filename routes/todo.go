package routes

import (
	"github.com/AnubhavPurohit691/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {

	// userGroup := r.Group("/todos")

	{
		r.POST("", controllers.TodoCreate)
		r.GET("", controllers.Todoall)
		r.GET("/:id", controllers.TodoShow)
	}
}
