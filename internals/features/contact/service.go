package contact

import "github.com/sub-rat/MyNewContactbook/internals/models"

type ServiceInterface interface {
	Query(offset, limit int, query string) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(req *Contact) (Contact, error)
	Update(id uint, update *Contact) (Contact, error)
	Delete(id uint) error
}

type service struct {
	repo RepositoryInterface
}

type Contact struct {
	models.Contact
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &service{repo}
}

func (service *service) Query(offset, limit int, query string) ([]Contact, error) {
	dataList, err := service.repo.Query(offset, limit, query)
	if err != nil {
		return []Contact{}, err
	}
	return dataList, nil
}

func (service *service) Get(id uint) (Contact, error) {
	contact, err := service.repo.Get(id)
	return contact, err
}

func (service *service) Create(req *Contact) (Contact, error) {
	//validation
	err := service.repo.Create(req)
	if err != nil {
		return Contact{}, err
	}
	return *req, nil
}

func (service *service) Update(id uint, update *Contact) (Contact, error) {
	err := service.repo.Update(id, update)
	if err != nil {
		return Contact{}, err
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
