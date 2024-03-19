package repository

import (
	"database/sql"
	"errors"
)

type User struct {
	UserID    string
	Hashedpwd string
	Email     string
	About     sql.NullString
	Name      sql.NullString
	Language  sql.NullString
	Picture   sql.NullString
}

func FindByCredentials(email string) (User, error) {
	query, err := connectedDb.Query("SELECT * FROM users WHERE email=$1", email)
	var user User
	if err != nil {
		return user, err

	}
	if query.Next() {
		err = query.Scan(&user.UserID, &user.Hashedpwd, &user.Email, &user.About, &user.Name, &user.Language, &user.Picture)
	}

	return user, err
}

func SignupUser(user User) error {

	if user.UserID == "" && user.Email == "" && user.Hashedpwd == "" {
		return errors.New("no userid, email or hashedpwd passed")
	}
	_, err := connectedDb.Exec("INSERT INTO users (userid, hashedpwd, email) VALUES ($1, $2, $3)", user.UserID, user.Hashedpwd, user.Email)

	return err
}
