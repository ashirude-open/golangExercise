package main

import (
	"github.com/gin-gonic/gin"
	"golangExercise/routes"
)

func main() {
	//r := gin.Default()
	r := gin.New()
	routes.UserRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
