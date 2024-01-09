package db

import (
	"context"
	"net/http"

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

func (d *Database) CreateSession(w http.ResponseWriter, r *http.Request, user *types.User) error {
	session, err := d.sessionStore.New(r, "user-session")
	if err != nil {
		return err
	}
	session.Values["userId"] = user.UserId
	session.Values["username"] = user.Username
	session.Values["role"] = user.Role
	session.Values["email"] = user.Email
	session.Values["authorized"] = true

	session.Save(r, w)

	return nil
}

func (d *Database) SerialiseSession(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	session, _ := d.sessionStore.Get(r, "user-session")
	ctx := context.WithValue(r.Context(), "user-id", session.Values["userId"])
	r = r.WithContext(ctx)
	return w, r, nil
}
