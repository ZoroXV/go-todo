package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks/create", postCreateTask)

	router.Run()
}

type task struct {
	ID		int		`json:"id"`
	Desc	string	`json:"desc"`
}

var tasks = []task{
	{ID: 0, Desc: "Task 0"},
	{ID: 1, Desc: "Task 1"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postCreateTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
