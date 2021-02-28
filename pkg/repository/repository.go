package repository

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user clientsCRUD.User) (int, error)
	GetUser(username, password string) (clientsCRUD.User, error)
}

type Client interface {
	Create(userId int, client clientsCRUD.Client) (int, error)
	GetAll(userId int) ([]clientsCRUD.Client, error)
	GetById(userId, clientId int) (clientsCRUD.Client, error)
	Delete(userId, clientId int) error
	Update(userId, clientId int, client clientsCRUD.Client) error
}

type Sale interface {
	Create(userId int, sale clientsCRUD.Sale) (int, error)
	GetAll(userId int) ([]clientsCRUD.Sale, error)
	GetById(userId, saleId int) (clientsCRUD.Sale, error)
	Delete(userId, saleId int) error
	Update(userId, saleId int, client clientsCRUD.Sale) error
}

type Repository struct {
	Authorization
	Client
	Sale
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		Client:        NewClientDB(db),
		Sale:          NewSaleDB(db),
	}
}
