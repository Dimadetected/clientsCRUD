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
func (s *ClientService) Create(client clientsCRUD.Client) (int, error) {
	return s.repo.Create(client)
}
func (s *ClientService) GetAll() ([]clientsCRUD.Client, error) {
	return s.repo.GetAll()
}
func (s *ClientService) GetById(clientId int) (clientsCRUD.Client, error) {
	return s.repo.GetById(clientId)
}
func (s *ClientService) Delete(clientId int) error {
	return s.repo.Delete(clientId)
}
func (s *ClientService) Update(clientId int, input clientsCRUD.UpdateClient) error {
	return s.repo.Update(clientId, input)
}
