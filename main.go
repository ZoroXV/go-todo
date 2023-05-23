package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

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

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.Run()
}
