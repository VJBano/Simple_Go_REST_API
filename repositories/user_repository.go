package repositories

import "vjbano.github.com/simple_crud_go/models"

type UserRepositoryImpl struct{}

func (u UserRepositoryImpl) Create(new models.User) (models.User, error) {

	//query from database here

	return models.User{
		Name: new.Name,
		Age:  new.Age,
	}, nil
}

func (u UserRepositoryImpl) GetAll() ([]models.User, error) {

	//query from database here

	return []models.User{
		{Name: "Doe", Age: 15},
	}, nil

}

func (u UserRepositoryImpl) GetByID(id string) (models.User, error) {

	//query from database here

	return models.User{
		Name: "Mitch",
		Age:  11,
	}, nil
}
func (u UserRepositoryImpl) Update(id string) (models.User, error) {

	return models.User{
		Name: "Mitch",
		Age:  11,
	}, nil
}
func (u UserRepositoryImpl) Delete(id string) (bool, error) {

	return true, nil
}

func (u UserRepositoryImpl) DeleteAll() (bool, error) {

	return true, nil
}
