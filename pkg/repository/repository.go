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
	Create(client clientsCRUD.Client) (int, error)
	GetAll() ([]clientsCRUD.Client, error)
	GetById(clientId int) (clientsCRUD.Client, error)
	Delete(clientId int) error
	Update(clientId int, input clientsCRUD.UpdateClient) error
}

type Sale interface {
	Create(userId int, sale clientsCRUD.Sale) (int, error)
	GetAll(userId int) ([]clientsCRUD.Sale, error)
	GetById(userId, saleId int) (clientsCRUD.Sale, error)
	Delete(userId, saleId int) error
	Update(userId, saleId int, input clientsCRUD.UpdateSale) error
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
