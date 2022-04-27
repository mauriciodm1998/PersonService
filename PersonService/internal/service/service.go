package service

import (
	"github.com/google/uuid"
)

var ()

type service struct {
	repo repositories.Repository
}

type Service interface {
	Create(u canonical.Person) (string, error)
	Get() ([]canonical.Person, error)
}

func New() Service {
	return &service{
		repo: repositories.New(),
	}
}

func (s *service) Create(u canonical.Person) (string, error) {

	if u.Id == "" {
		u.Id = uuid.NewString()
	}

	if err := s.repo.Create(u); err != nil {
		return "", err
	}

	return u.Id, nil
}

func (s *service) Get() ([]canonical.Person, error) {
	return s.repo.Get()
}
