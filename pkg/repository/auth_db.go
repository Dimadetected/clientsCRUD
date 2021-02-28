package repository

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{
		db: db,
	}
}

func (r *AuthDB) CreateUser(user clientsCRUD.User) (int, error) {
	return 0, nil
}
func (r *AuthDB) GetUser(username, password string) (clientsCRUD.User, error) {
	return clientsCRUD.User{}, nil

}
