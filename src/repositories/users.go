package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (u users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, error := u.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if error != nil {
		return nil, error
	}


	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) SearchByID(ID uint64) (models.User, error) {
	lines, error := u.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		ID,
	)
	if error != nil {
		return models.User{}, error
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (u users) Update(ID uint64, user models.User) error {
	statement, error := u.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if error != nil {
		return error
	}
	
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}

func (u users) Delete(ID uint64) error {
	statement, error := u.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

func (u users) SearchByEmail(email string) (models.User, error) {
	line, error := u.db.Query("select id, password from users where email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if error = line.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
 }

func (u users) Follow(userToFollowId, followerId uint64) error {
	statement, error := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if error !=nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(userToFollowId, followerId); error != nil {
		return error
	}

	return nil
}