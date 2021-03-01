package repository

import (
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/jmoiron/sqlx"
)

type AuthDB struct {
	db *sqlx.DB
}

func NewAuthDB(db *sqlx.DB) *AuthDB {
	return &AuthDB{
		db: db,
	}
}

func (r *AuthDB) CreateUser(user clientsCRUD.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, username, password) VALUES($1, $2, $3) RETURNING id`, usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthDB) GetUser(username, password string) (clientsCRUD.User, error) {
	var user clientsCRUD.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE username = $1 and password = $2`, usersTable)
	err := r.db.Get(&user, query, username, password)

	if err != nil {
		return clientsCRUD.User{}, err
	}

	return user, err

}
