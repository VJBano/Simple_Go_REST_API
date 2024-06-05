package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vjbano.github.com/simple_crud_go/interfaces"
	"vjbano.github.com/simple_crud_go/models"
)

type UserControllers struct {
	Service interfaces.UserService
}

func (s UserControllers) CreateNewUser(ctx *gin.Context) {

	newUser, err := s.Service.CreateNewUser(models.User{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (s UserControllers) GetAllUser(ctx *gin.Context) {

	users, err := s.Service.GetAllUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (s UserControllers) GetUserById(ctx *gin.Context) {

	user, err := s.Service.DeleteUserById("sample")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s UserControllers) UpdateUser(ctx *gin.Context) {

	user, err := s.Service.UpdateUser("sample")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s UserControllers) DeleteUserByID(ctx *gin.Context) {

	user, err := s.Service.GetUserById("sample")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (s UserControllers) DeleteAll(ctx *gin.Context) {

	user, err := s.Service.DeleteAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
