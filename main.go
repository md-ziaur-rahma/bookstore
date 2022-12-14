package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
