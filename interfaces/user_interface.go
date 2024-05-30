package interfaces

import "vjbano.github.com/simple_crud_go/models"

type UserRepository interface {
	Create(newUser models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetByID(userID string) (models.User, error)
	Update(userID string) (models.User, error)
	Delete(userID string) (bool, error)
	DeleteAll() (bool, error)
}

type UserService interface {
	CreateNewUser(models.User) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserById(userID string) (models.User, error)
	UpdateUser(userID string) (models.User, error)
	DeleteUserById(userID string) (bool, error)
	DeleteAll() (bool, error)
}
