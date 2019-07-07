package main

import (
	"github.com/yukinagae/go-api-template/src/todo"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Aloha, World!",})
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		todoID := c.Param("id")
		t, err := todo.GetTodo(todoID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return 
		}
		c.JSON(200, t)
	})

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todo.GetTodos())
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
