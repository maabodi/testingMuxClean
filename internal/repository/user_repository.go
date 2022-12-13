package repository

import (
	"cobaclean/internal/domain"
	"cobaclean/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type userMysql struct {
	DB *gorm.DB
}

func (um *userMysql) GetAllUsers() []model.User {
	res := []model.User{}
	um.DB.Find(&res)

	return res
}

func (um *userMysql) GetOneByID(id int) (user model.User, err error) {
	res := um.DB.Where("users_id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func NewUserMysql(db *gorm.DB) domain.UserAdapterRepository {
	return &userMysql{
		DB: db,
	}
}
