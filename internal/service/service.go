package service

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/repository"
)

type Services interface {
	CreateContact(contact someRestApi.Contacts) (int64, error)
	GetContact(id int64) (someRestApi.Contacts, error)
	UpdateContact(id int64, contact someRestApi.Contacts) (someRestApi.Contacts, error)
	DeleteContact(id int64) (int64, error)
}

type Authorization interface {
}

type Service struct {
	Services
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Services: NewContactService(repos),
	}
}
