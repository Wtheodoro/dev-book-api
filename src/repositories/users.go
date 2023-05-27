package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, error := u.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")
	if error != nil {
		return 0, nil
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, nil
	}

	lastInsertID, error := result.LastInsertId()
	if error != nil {
		return 0, nil
	}

	return uint64(lastInsertID), nil
}