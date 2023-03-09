package service

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/repository"
)

type ContactService struct {
	contact repository.Contact
}

func NewContactService(contact repository.Contact) *ContactService {
	return &ContactService{contact: contact}
}

func (s *ContactService) CreateContact(contact someRestApi.Contacts) (int64, error) {
	return s.contact.CreateContact(contact)
}

func (s *ContactService) GetContact(id int64) (someRestApi.Contacts, error) {
	return s.contact.GetContact(id)
}

func (s *ContactService) UpdateContact(id int64, contact someRestApi.Contacts) (someRestApi.Contacts, error) {
	return s.contact.UpdateContact(id, contact)
}

func (s *ContactService) DeleteContact(id int64) (int64, error) {
	return s.contact.DeleteContact(id)
}
