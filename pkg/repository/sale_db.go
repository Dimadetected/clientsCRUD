package repository

import (
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SaleDB struct {
	db *sqlx.DB
}

func NewSaleDB(db *sqlx.DB) *SaleDB {
	return &SaleDB{
		db: db,
	}
}

func (r *SaleDB) GetAll(clientId int) ([]clientsCRUD.Sale, error) {
	var sales []clientsCRUD.Sale
	query := fmt.Sprintf(`SELECT sl.id, sl.total, sl.count FROM %s sl INNER JOIN %s cs on sl.id = cs.sale_id INNER JOIN %s cl on cs.client_id = cl.id WHERE cl.id = $1`, salesTable, clientsSalesTable, clientsTable)
	if err := r.db.Select(&sales, query, clientId); err != nil {
		return []clientsCRUD.Sale{}, err
	}

	return sales, nil
}
func (r *SaleDB) GetById(saleId int) (clientsCRUD.Sale, error) {
	var sale clientsCRUD.Sale
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, salesTable)
	if err := r.db.Get(&sale, query, saleId); err != nil {
		return clientsCRUD.Sale{}, err
	}

	return sale, nil
}

func (r *SaleDB) Create(clientId int, sale clientsCRUD.Sale) (int, error) {
	txx, err := r.db.Begin()
	var saleId int
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf(`INSERT INTO %s (count, total) VALUES ($1, $2) RETURNING id`, salesTable)
	row := txx.QueryRow(query, sale.Count, sale.Total)
	if err := row.Scan(&saleId); err != nil {
		txx.Rollback()
		return 0, err
	}

	queryClientsSale := fmt.Sprintf(`INSERT INTO %s (client_id, sale_id) VALUES ($1, $2)`, clientsSalesTable)
	_, err = txx.Exec(queryClientsSale, clientId, saleId)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	txx.Commit()
	return saleId, nil
}

func (r *SaleDB) Delete(saleId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, salesTable)
	_, err := r.db.Exec(query, saleId)
	return err
}
func (r *SaleDB) Update(saleId int, input clientsCRUD.UpdateSale) error {
	var setValues = make([]string, 0)
	if input.Total != nil {
		setValues = append(setValues, fmt.Sprintf("total=%f", *input.Total))
	}
	if input.Count != nil {
		setValues = append(setValues, fmt.Sprintf("count=%d", *input.Count))
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $1`, salesTable, setQuery)
	_, err := r.db.Exec(query, saleId)
	return err
}
