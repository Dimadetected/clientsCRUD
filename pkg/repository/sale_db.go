package repository

import (
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
)

type SaleDB struct {
	db *sqlx.DB
}

func NewSaleDB(db *sqlx.DB) *SaleDB {
	return &SaleDB{
		db: db,
	}
}

func (r *SaleDB) GetAll(userId int) ([]clientsCRUD.Sale, error) {
	return []clientsCRUD.Sale{}, nil
}
func (r *SaleDB) GetById(userId, saleId int) (clientsCRUD.Sale, error) {
	return clientsCRUD.Sale{}, nil
}

func (r *SaleDB) Create(userId int, sale clientsCRUD.Sale) (int, error) {
	return 0, nil
}

func (r *SaleDB) Delete(userId, saleId int) error {
	return nil
}
func (r *SaleDB) Update(userId, saleId int, sale clientsCRUD.Sale) error {
	return nil
}
