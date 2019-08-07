package user

import (
	"crypto/sha256"
	"fmt"

	"wume-composer/internal/pkg/common/verifier"
	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/models"
)

func PasswordHash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func SignIn(signInData models.SignInData) (jwt models.JwtData, err error, incorrectFields []string) {
	var userData db.UserData
	passHash := PasswordHash(signInData.Password)

	if verifier.IsEmail(signInData.Login) {
		userData, err = db.UserGetByEmailAndPassHash(signInData.Login, passHash)
		if err != nil {
			if err == models.NotFoundError {
				return jwt, models.IncorrectDataError, []string{"password"}
			}
			return
		}
	} else if verifier.IsUsername(signInData.Login) {
		userData, err = db.UserGetByUsernameAndPassHash(signInData.Login, passHash)
		if err != nil {
			if err == models.NotFoundError {
				return jwt, models.IncorrectDataError, []string{"password"}
			}
			return
		}
	} else {
		return jwt, models.IncorrectDataError, []string{"login"}
	}

	return models.JwtData{
		Id:       userData.Id,
		Email:    userData.Email,
		Username: userData.Username,
	}, nil, nil
}

func GetUser(username string) (models.UserData, error) {
	userData, err := db.UserGetByUsername(username)
	if err != nil {
		return models.UserData{}, err
	}

	return models.UserData{
		Id:       userData.Id,
		Email:    userData.Email,
		Username: userData.Username,
		Avatar: userData.Avatar,
	}, nil
}

func CreateUser(signUpData models.SignUpData) (jwt models.JwtData, err error, incorrectFields []string) {
	// Check email uniqueness
	id, err := db.UserGetIdByEmail(signUpData.Email)
	if err != nil {
		return jwt, err, nil
	}
	if id != 0 {
		incorrectFields = append(incorrectFields, "email")
	}

	// Check username uniqueness
	id, err = db.UserGetIdByUsername(signUpData.Username)
	if err != nil {
		return jwt, err, nil
	}
	if id != 0 {
		incorrectFields = append(incorrectFields, "username")
	}

	// If something is already in use
	if len(incorrectFields) > 0 {
		return jwt, models.AlreadyExistsError, incorrectFields
	}

	// Create user
	id, err = db.UserCreate(db.UserData{
		Email:    signUpData.Email,
		Username: signUpData.Username,
		Password: PasswordHash(signUpData.Password),
	})
	if err != nil {
		return jwt, err, nil
	}
	return models.JwtData{
		Id:       id,
		Email:    signUpData.Email,
		Username: signUpData.Username,
	}, nil, nil
}

func UpdateUser(id int64, userData models.UpdateUserData) (jwt models.JwtData, err error, incorrectFields []string) {
	// Check email uniqueness
	usedId, err := db.UserGetIdByEmail(userData.Email)
	if err != nil {
		return jwt, err, nil
	}
	if usedId != 0 {
		incorrectFields = append(incorrectFields, "email")
	}

	// Check username uniqueness
	usedId, err = db.UserGetIdByUsername(userData.Username)
	if err != nil {
		return jwt, err, nil
	}
	if usedId != 0 {
		incorrectFields = append(incorrectFields, "username")
	}

	// If something is already in use
	if len(incorrectFields) > 0 {
		return jwt, models.AlreadyExistsError, incorrectFields
	}

	// Update user
	err = db.UserUpdateData(db.UserData{
		Id:       id,
		Email:    userData.Email,
		Username: userData.Username,
	})
	if err != nil {
		return models.JwtData{}, err, nil
	}

	return models.JwtData{
		Id:       id,
		Email:    userData.Email,
		Username: userData.Username,
	}, nil, nil
}

func UpdatePassword(id int64, passwordData models.UpdatePasswordData) (error, []string) {
	err := db.UserUpdatePassword(id, PasswordHash(passwordData.OldPassword), PasswordHash(passwordData.NewPassword))
	if err == models.IncorrectDataError {
		return models.IncorrectDataError, []string{"password"}
	}
	return err, nil
}

func RemoveUser(id int64, removeData models.RemoveUserData) (error, []string) {
	err := db.UserRemove(id, PasswordHash(removeData.Password))
	if err == models.IncorrectDataError {
		return models.IncorrectDataError, []string{"password"}
	}
	return err, nil
}

func GetUsers(page uint, limit uint) (models.UsersData, error) {
	usersData, count, err := db.UsersGet(limit, (page-1)*limit)
	if err != nil {
		return models.UsersData{}, err
	}

	return models.UsersData{
		Users: usersData,
		Count: count,
	}, nil
}
