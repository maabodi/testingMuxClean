package service

import (
	"cobaclean/internal/config"
	"cobaclean/internal/domain"
	"cobaclean/internal/model"
)

type sUser struct {
	c    config.Config
	repo domain.UserAdapterRepository
}

func (s *sUser) GetAllUsersService() []model.User {
	return s.repo.GetAllUsers()
}

func (s *sUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneByID(id)
}

func NewServiceUser(repo domain.UserAdapterRepository, c config.Config) domain.UserAdapterService {
	return &sUser{
		repo: repo,
		c:    c,
	}
}
