package service

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
)

type ClientService struct {
	repo repository.Client
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{
		repo: repo,
	}
}
func (s *ClientService) Create(userId int, client clientsCRUD.Client) (int, error) {
	return s.repo.Create(userId, client)
}
func (s *ClientService) GetAll(userId int) ([]clientsCRUD.Client, error) {
	return s.repo.GetAll(userId)
}
func (s *ClientService) GetById(userId, clientId int) (clientsCRUD.Client, error) {
	return s.repo.GetById(userId, clientId)
}
func (s *ClientService) Delete(userId, clientId int) error {
	return s.repo.Delete(userId, clientId)
}
func (s *ClientService) Update(userId, clientId int, client clientsCRUD.Client) error {
	return s.repo.Update(userId, clientId, client)
}
