package controllers

import "vjbano.github.com/simple_crud_go/models"

type UserModels struct {
	*models.User
}

func (p *UserModels) GetName() string {

	return p.Name
}

func (p *UserModels) GetAge() int {
	return p.Age
}
