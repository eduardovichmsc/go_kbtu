package usecase

import (
	"practice_3/internal/pkg/modules"
	"practice_3/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetAll() ([]modules.User, error) {
	return u.repo.GetUsers()
}

func (u *UserUsecase) GetByID(id int) (*modules.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) Create(user *modules.User) (int, error) {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) Update(id int, user *modules.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *UserUsecase) Delete(id int) error {
	return u.repo.DeleteUser(id)
}