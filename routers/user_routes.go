package routers

import (
	"github.com/gin-gonic/gin"
	"vjbano.github.com/simple_crud_go/controllers"
	"vjbano.github.com/simple_crud_go/repositories"
	"vjbano.github.com/simple_crud_go/services"
)

func UserRoutes(in *gin.Engine) {

	userRepository := repositories.UserRepositoryImpl{}
	userService := services.UserServiceImpl{Repository: userRepository}
	userControllers := controllers.UserControllers{Service: userService}

	in.POST("/api/user/create", userControllers.CreateNewUser)
	in.GET("/api/user/users", userControllers.GetAllUser)
}
