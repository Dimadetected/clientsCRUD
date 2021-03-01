package repository

import (
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ClientDB struct {
	db *sqlx.DB
}

func NewClientDB(db *sqlx.DB) *ClientDB {
	return &ClientDB{
		db: db,
	}
}

func (r *ClientDB) GetAll() ([]clientsCRUD.Client, error) {
	var clients []clientsCRUD.Client
	query := fmt.Sprintf(`SELECT * FROM %s`, clientsTable)
	if err := r.db.Select(&clients, query); err != nil {
		return nil, err
	}

	return clients, nil
}
func (r *ClientDB) GetById(clientId int) (clientsCRUD.Client, error) {
	var client clientsCRUD.Client
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, clientsTable)

	if err := r.db.Get(&client, query, clientId); err != nil {
		return clientsCRUD.Client{}, err
	}

	return client, nil
}
func (r *ClientDB) Create(client clientsCRUD.Client) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, email, age) VALUES ($1, $2, $3) RETURNING id`, clientsTable)
	row := r.db.QueryRow(query, client.Name, client.Email, client.Age)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *ClientDB) Delete(clientId int) error {

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, clientsTable)

	_, err := r.db.Exec(query, clientId)

	return err
}
func (r *ClientDB) Update(clientId int, input clientsCRUD.UpdateClient) error {
	setValues := make([]string, 0)
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name='%s'", *input.Name))
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email='%s'", *input.Email))
	}
	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=%d", *input.Age))
	}

	setQuery := strings.Join(setValues, ",")
	fmt.Println(setQuery)
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = %d`, clientsTable, setQuery, clientId)
	_, err := r.db.Exec(query)
	return err
}
