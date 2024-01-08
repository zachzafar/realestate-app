package types

type User struct {
	UserId        int
	Username      string
	Password_hash string
	Role          string
	Email         string
}

func NewUser(userId int, username string, password string, Role string, Email string) *User {
	return &User{UserId: userId, Username: username, Password_hash: password, Role: Role, Email: Email}
}
