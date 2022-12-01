package main

import (
	"golangExercise/config"
	"golangExercise/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	config.Connect()
	routes.UserRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
