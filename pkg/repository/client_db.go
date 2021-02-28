package repository

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
)

type ClientDB struct {
	db *sqlx.DB
}

func NewClientDB(db *sqlx.DB) *ClientDB {
	return &ClientDB{
		db: db,
	}
}

func (r *ClientDB) GetAll(userId int) ([]clientsCRUD.Client, error) {
	return []clientsCRUD.Client{}, nil
}
func (r *ClientDB) GetById(userId, clientId int) (clientsCRUD.Client, error) {
	return clientsCRUD.Client{}, nil
}
func (r *ClientDB) Create(userId int, client clientsCRUD.Client) (int, error) {
	return 0, nil
}
func (r *ClientDB) Delete(userId, clientId int) error {
	return nil
}
func (r *ClientDB) Update(userId, clientId int, client clientsCRUD.Client) error {
	return nil
}
