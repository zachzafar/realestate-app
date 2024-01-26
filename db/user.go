package db

import (
	"openlettings.com/types"
)

func (d *Database) GetUser(email string) (*types.User, error) {
	var userId int
	var password_hash string

	query := `SELECT user_id,password_hash
				FROM users
				WHERE email = $1`

	err := d.db.QueryRow(query, email).Scan(&userId, &password_hash)

	if err != nil {
		return nil, err
	}

	return &types.User{UserId: userId, PasswordHash: password_hash}, nil

}

func (d *Database) CreateUser(user *types.User) error {
	query := `INSERT INTO users (password_hash, email) VALUES ($1,$2)`
	_, err := d.db.Exec(query, user.PasswordHash, user.Email)
	return err
}
