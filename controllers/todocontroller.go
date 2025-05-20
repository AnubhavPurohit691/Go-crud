package controllers

import (
	"github.com/AnubhavPurohit691/go-crud/initializer"
	"github.com/AnubhavPurohit691/go-crud/models"
	"github.com/gin-gonic/gin"
)

func TodoCreate(c *gin.Context) {
	var body struct {
		Content string
		Status  bool
	}

	c.Bind(&body)

	todo := models.Todo{Content: body.Content, Status: body.Status}
	result := initializer.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodoShow(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	initializer.DB.First(&todo, id)
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func Todoall(c *gin.Context) {
	var todos []models.Todo
	initializer.DB.Find(&todos)
	c.JSON(200, gin.H{
		"todos": todos,
	})
}
