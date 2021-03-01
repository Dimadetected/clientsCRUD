package service

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
)

type SaleService struct {
	repo repository.Sale
}

func NewSaleService(repo repository.Sale) *SaleService {
	return &SaleService{
		repo: repo,
	}
}
func (s *SaleService) Create(userId int, client clientsCRUD.Sale) (int, error) {
	return s.repo.Create(userId, client)
}
func (s *SaleService) GetAll(userId int) ([]clientsCRUD.Sale, error) {
	return s.repo.GetAll(userId)
}
func (s *SaleService) GetById(userId, clientId int) (clientsCRUD.Sale, error) {
	return s.repo.GetById(userId, clientId)
}
func (s *SaleService) Delete(userId, clientId int) error {
	return s.repo.Delete(userId, clientId)
}
func (s *SaleService) Update(userId, clientId int, input clientsCRUD.UpdateSale) error {
	return s.repo.Update(userId, clientId, input)
}
