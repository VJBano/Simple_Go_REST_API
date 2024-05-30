package main

import (
	"github.com/gin-gonic/gin"
	"vjbano.github.com/simple_crud_go/database"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger())

	database.DBConnection()
	router.Run(":4040")

}
