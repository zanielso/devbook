package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersReposiotry(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nick, email, password) value (?, ?, ?, ?)",
	)
	if error != nil {
		return 0, nil
	}
	defer statement.Close()
	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, nil
	}

	lastInsertId, error := result.LastInsertId()
	if error != nil {
		return 0, nil
	}

	return uint64(lastInsertId), nil

}

func (repository Users) FindById(userId uint) (models.User, error) {
	rows, error := repository.db.Query("select id, name, nick , email, createdAt  from users u where id = ?", userId)
	if error != nil {
		return models.User{}, error
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if error = rows.Scan(
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
