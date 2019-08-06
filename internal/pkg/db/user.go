package db

import (
	"database/sql"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"wume-composer/internal/pkg/models"
)

type UserData struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"pass_hash"`
}

func UserGetIdByEmail(email string) (id int64, err error) {
	id, err = isExists("users", "email = $1", email)
	if err != nil {
		return
	}
	return id, nil
}

func UserGetIdByUsername(username string) (id int64, err error) {
	id, err = isExists("users", "username = $1", username)
	if err != nil {
		return
	}
	return id, nil
}

func UserCreate(data UserData) (id int64, err error) {
	id, err = isExists("users", "email = $1 OR username = $2",
		data.Email, data.Username)
	if err != nil {
		return
	}
	if id != 0 {
		return id, models.AlreadyExistsError
	}

	return insert("users", "username, email, pass_hash", "$1, $2, $3",
		data.Username, data.Email, data.Password)
}

func UserGetByUsername(username string) (data UserData, err error) {
	row, err := findRowBy("users", "id, username, email, pass_hash", "username = $1", username)
	if err != nil {
		return
	}
	err = row.Scan(&data.Id, &data.Username, &data.Email, &data.Password)
	if err == sql.ErrNoRows {
		return data, models.NotFoundError
	}
	return
}

func UserGetByEmailAndPassHash(email string, passHash string) (data UserData, err error) {
	row, err := findRowBy("users", "id, username, email, pass_hash",
		"email = $1 AND pass_hash = $2", email, passHash)
	if err != nil {
		return
	}
	err = row.Scan(&data.Id, &data.Username, &data.Email, &data.Password)
	if err == sql.ErrNoRows {
		return data, models.NotFoundError
	}
	return
}

func UserGetByUsernameAndPassHash(username string, passHash string) (data UserData, err error) {
	row, err := findRowBy("users", "id, username, email, pass_hash",
		"username = $1 AND pass_hash = $2", username, passHash)
	if err != nil {
		return
	}
	err = row.Scan(&data.Id, &data.Username, &data.Email, &data.Password)
	if err == sql.ErrNoRows {
		return data, models.NotFoundError
	}
	return
}

func UserUpdateData(data UserData) error {
	id, err := isExists("users", "id = $1", data.Id)
	if err != nil {
		return err
	}
	if id == 0 {
		return models.NotFoundError
	}

	_, err = updateBy("users", "username = $1, email = $2", "id = $3",
		data.Username, data.Email, data.Id)
	if sqlError, ok := err.(*pq.Error); ok {
		// Unique violation
		if sqlError.Code == uniqueErrorCode {
			return models.AlreadyExistsError
		}
	}
	return err
}

func UserCheckPassword(id int64, passHash string) error {
	row, err := findRowBy("users", "id, pass_hash", "id = $1", id)
	if err != nil {
		return err
	}
	var realPassHash string
	err = row.Scan(&id, &realPassHash)
	if err == sql.ErrNoRows {
		return models.NotFoundError
	}

	if passHash != realPassHash {
		return models.IncorrectDataError
	}
	return nil
}

func UserUpdatePassword(id int64, oldPassHash string, newPassHash string) (err error) {
	if err = UserCheckPassword(id, oldPassHash); err != nil {
		return err
	}

	_, err = updateBy("users", "pass_hash = $1", "id = $2", newPassHash, id)
	return err
}

func UserRemove(id int64, passHash string) error {
	if err := UserCheckPassword(id, passHash); err != nil {
		return err
	}

	_, err := removeBy("users", "id = $1", id)
	return err
}

func UserTruncate() error {
	return truncate("users")
}
