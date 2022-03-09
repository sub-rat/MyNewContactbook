package user

import (
	"errors"
	"github.com/sub-rat/MyNewContactbook/internals/models"
	"github.com/sub-rat/MyNewContactbook/pkg/utils"
)

type ServiceInterface interface {
	Query(offset, limit int, query string, fieldType string) ([]User, error)
	Get(id uint) (User, error)
	Create(req *User) (User, error)
	Update(id uint, update *User) (User, error)
	Delete(id uint) error
}

type service struct {
	repo RepositoryInterface
}

type User struct {
	models.User
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &service{repo}
}

func (service *service) Query(offset, limit int, query string, fieldType string) ([]User, error) {
	dataList, err := service.repo.Query(offset, limit, query, fieldType)
	if err != nil {
		return []User{}, err
	}
	return dataList, nil
}

func (service *service) Get(id uint) (User, error) {
	contact, err := service.repo.Get(id)
	return contact, err
}

func (service *service) Create(req *User) (User, error) {
	// Todo Validate the data, name validation, email validation, password length, password strength
	if req.FullName == "" {
		return User{}, errors.New(" Fullname required ")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return User{}, err
	}
	req.Password = hashedPassword
	err = service.repo.Create(req)
	if err != nil {
		return User{}, err
	}
	return *req, nil
}

func (service *service) Update(id uint, update *User) (User, error) {
	// ToDo Check for only updated value
	// create map of updated value and send
	err := service.repo.Update(id, update)
	if err != nil {
		return User{}, err
	}
	return *update, nil
}

func (service *service) Delete(id uint) error {
	err := service.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
