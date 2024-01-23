package db

import (
	"openlettings.com/types"
)

func (d *Database) GetUser(email string) (*types.User, error) {
	var username string
	var userId int
	var password_hash string
	var role string
	query := `SELECT user_id,username,password_hash,role
				FROM users
				WHERE email = ?;`

	err := d.db.QueryRow(query, email).Scan(&userId, &username, &password_hash, &role)

	if err != nil {
		return nil, err
	}

	return types.NewUser(userId, username, password_hash, role, email), nil

}

func (d *Database) CreateUser(user *types.User) error {
	query := `INSERT INTO users (username,password_hash,role, email) VALUES ($1,$2,$3,$4)`
	_, err := d.db.Exec(query, user.Username, user.Password_hash, user.Role, user.Email)

	return err
}
