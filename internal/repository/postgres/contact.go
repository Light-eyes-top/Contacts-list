package postgres

import (
	someRestApi "SomeRestApi"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ContactRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) CreateContact(contact someRestApi.Contacts) (int64, error) {
	var id int64
	line := r.db.QueryRow("INSERT INTO contacts (fio, email, phone) VALUES ($1, $2, $3) RETURNING id",
		contact.Fio, contact.Email, contact.Phone)
	if err := line.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ContactRepository) GetContact(id int64) (someRestApi.Contacts, error) {
	var line someRestApi.Contacts
	query := fmt.Sprintf("SELECT * FROM contacts WHERE id = $1")
	err := r.db.Get(&line, query, id)

	return line, err
}

func (r *ContactRepository) UpdateContact(id int64, contact someRestApi.Contacts) (someRestApi.Contacts, error) {
	query := fmt.Sprintf("UPDATE contacts SET fio = $1, email = $2, phone = $3 WHERE id = $4")
	_, err := r.db.Exec(query, contact.Fio, contact.Email, contact.Phone, id)

	result := someRestApi.Contacts{
		Id:    id,
		Fio:   contact.Fio,
		Email: contact.Email,
		Phone: contact.Phone,
	}

	return result, err
}

func (r *ContactRepository) DeleteContact(id int64) (int64, error) {
	query := fmt.Sprintf("DELETE FROM contacts WHERE id = $1")
	_, err := r.db.Exec(query, id)

	return id, err
}
