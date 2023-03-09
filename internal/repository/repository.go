package repository

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/repository/in_memory"
	"SomeRestApi/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Contact interface {
	CreateContact(contact someRestApi.Contacts) (int64, error)
	GetContact(id int64) (someRestApi.Contacts, error)
	UpdateContact(id int64, contact someRestApi.Contacts) (someRestApi.Contacts, error)
	DeleteContact(id int64) (int64, error)
}

type Authorization interface {
}

type Repository struct {
	Contact
	Authorization
}

func NewRepositoryInMemory() *Repository {
	return &Repository{
		Contact: in_memory.NewContactRepository(),
	}
}

func NewRepositoryPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		Contact: postgres.NewContactRepository(db),
	}
}
