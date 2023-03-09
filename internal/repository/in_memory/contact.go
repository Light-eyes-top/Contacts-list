package in_memory

import (
	someRestApi "SomeRestApi"
	"sync"
)

type ContactRepository struct {
	db map[int64]someRestApi.Contacts
	id int64
	mu *sync.RWMutex
}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{db: make(map[int64]someRestApi.Contacts), id: 1, mu: &sync.RWMutex{}}
}

func (r *ContactRepository) CreateContact(contact someRestApi.Contacts) (int64, error) {
	r.mu.Lock()
	r.db[r.id] = someRestApi.Contacts{
		Id:    r.id,
		Fio:   contact.Fio,
		Email: contact.Email,
		Phone: contact.Phone,
	}
	r.mu.Unlock()
	r.id++
	return r.id - 1, nil
}

func (r *ContactRepository) GetContact(id int64) (someRestApi.Contacts, error) {
	defer r.mu.RUnlock()
	r.mu.RLock()
	return r.db[id], nil
}

func (r *ContactRepository) UpdateContact(id int64, contact someRestApi.Contacts) (someRestApi.Contacts, error) {
	r.mu.Lock()
	r.db[id] = someRestApi.Contacts{
		Id:    id,
		Fio:   contact.Fio,
		Email: contact.Email,
		Phone: contact.Phone,
	}
	r.mu.Unlock()
	return r.db[id], nil
}

func (r *ContactRepository) DeleteContact(id int64) (int64, error) {
	r.mu.Lock()
	delete(r.db, id)
	r.mu.Unlock()
	return id, nil
}
