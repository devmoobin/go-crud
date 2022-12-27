package main

import (
	"devmoobin/go-crud/controllers"
	"devmoobin/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts",controllers.PostCreate)
	r.GET("/posts",controllers.PostsAll)
	r.GET("/post/:id",controllers.PostByIndex)
	r.PUT("/post/:id",controllers.PostUpdate)
	r.DELETE("/post/:id",controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}