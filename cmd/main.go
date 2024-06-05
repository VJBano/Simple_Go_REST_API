package main

import (
	"github.com/gin-gonic/gin"
	"vjbano.github.com/simple_crud_go/database"
	"vjbano.github.com/simple_crud_go/routers"
)

func main() {

	router := gin.Default()

	router.Use(gin.Logger())

	routers.UserRoutes(router)

	database.DBConnection()
	router.Run(":4040")

}
