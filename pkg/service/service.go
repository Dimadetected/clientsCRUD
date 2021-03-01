package service

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
)

type Authorization interface {
	CreateUser(user clientsCRUD.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Client interface {
	Create(client clientsCRUD.Client) (int, error)
	GetAll() ([]clientsCRUD.Client, error)
	GetById(clientId int) (clientsCRUD.Client, error)
	Delete(clientId int) error
	Update(clientId int, input clientsCRUD.UpdateClient) error
}

type Sale interface {
	Create(clientId int, sale clientsCRUD.Sale) (int, error)
	GetAll(clientId int) ([]clientsCRUD.Sale, error)
	GetById(saleId int) (clientsCRUD.Sale, error)
	Delete(saleId int) error
	Update(saleId int, input clientsCRUD.UpdateSale) error
}

type Service struct {
	Authorization
	Client
	Sale
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Client:        NewClientService(repo.Client),
		Sale:          NewSaleService(repo.Sale),
	}
}
