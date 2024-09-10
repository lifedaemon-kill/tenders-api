package repository

import "tenders-api/model"

type userRepository interface {
	GetUsers() []model.User
	GetUser() model.User
	CreateUser(model.User) model.User
	UpdateUser(model.User) model.User
	DeleteUser(model.User) model.User
}
