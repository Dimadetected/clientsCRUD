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
func (s *SaleService) Create(clientId int, client clientsCRUD.Sale) (int, error) {
	return s.repo.Create(clientId, client)
}
func (s *SaleService) GetAll(clientId int) ([]clientsCRUD.Sale, error) {
	return s.repo.GetAll(clientId)
}
func (s *SaleService) GetById(saleId int) (clientsCRUD.Sale, error) {
	return s.repo.GetById(saleId)
}
func (s *SaleService) Delete(saleId int) error {
	return s.repo.Delete(saleId)
}
func (s *SaleService) Update(saleId int, input clientsCRUD.UpdateSale) error {
	return s.repo.Update(saleId, input)
}
