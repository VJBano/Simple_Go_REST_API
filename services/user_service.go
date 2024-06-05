package services

import (
	"vjbano.github.com/simple_crud_go/interfaces"
	"vjbano.github.com/simple_crud_go/models"
)

type UserServiceImpl struct {
	Repository interfaces.UserRepository
}

func (u UserServiceImpl) CreateNewUser(user models.User) (models.User, error) {

	return u.Repository.Create(user)
}

func (u UserServiceImpl) GetAllUser() ([]models.User, error) {

	return u.Repository.GetAll()
}

func (u UserServiceImpl) GetUserById(id string) (models.User, error) {

	return u.Repository.GetByID(id)
}

func (u UserServiceImpl) UpdateUser(id string) (models.User, error) {

	return u.Repository.Update(id)
}

func (u UserServiceImpl) DeleteUserById(id string) (bool, error) {

	return u.Repository.Delete(id)
}

func (u UserServiceImpl) DeleteAll() (bool, error) {

	return u.Repository.DeleteAll()
}
