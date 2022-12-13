package domain

import "cobaclean/internal/model"

type UserAdapterRepository interface {
	// CreateUsers(user model.User) error
	GetAllUsers() []model.User
	GetOneByID(id int) (user model.User, err error)
	// GetOneByEmail(email string) (user model.User, err error)
	// UpdateOneByID(id int, user model.User) error
	// DeleteByID(id int) error
}

type UserAdapterService interface {
	// CreateUserService(user model.User) (error, int)
	// UpdateUserService(id int, user model.User) error
	GetAllUsersService() []model.User
	GetUserByID(id int) (model.User, error)
	// LoginUser(email, password string) (string, int)
	// DeleteByID(id int) error
}
