package auth

import (
	"crypto/sha256"
	"fmt"

	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/models"
	"wume-composer/internal/pkg/verifier"
)

func PasswordHash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func SignUp(signUpData models.SignUpData) (jwt models.JwtData, err error, incorrectFields []string) {
	// Check email uniqueness
	id, err := db.AuthGetIdByEmail(signUpData.Email)
	if err != nil {
		return jwt, err, nil
	}
	if id != 0 {
		incorrectFields = append(incorrectFields, "email")
	}

	// Check username uniqueness
	id, err = db.AuthGetIdByUsername(signUpData.Username)
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
	id, err = db.AuthCreate(db.AuthData{
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

func SignIn(signInData models.SignInData) (jwt models.JwtData, err error, incorrectFields []string) {
	var userData db.AuthData
	passHash := PasswordHash(signInData.Password)

	if verifier.IsEmail(signInData.Login) {
		userData, err = db.AuthFindByEmailAndPassHash(signInData.Login, passHash)
		if err != nil {
			if err == models.NotFoundError {
				return jwt, models.IncorrectDataError, []string{"password"}
			}
			return
		}
	} else if verifier.IsUsername(signInData.Login) {
		userData, err = db.AuthFindByUsernameAndPassHash(signInData.Login, passHash)
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

func UpdateAuth(id int64, userData models.UpdateUserData) (models.JwtData, error, []string) {
	err := db.AuthUpdateData(db.AuthData{
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
	err := db.AuthUpdatePassword(id, PasswordHash(passwordData.NewPassword))
	if err == models.IncorrectDataError {
		return models.IncorrectDataError, []string{"password"}
	}
	return err, nil
}

func RemoveAuth(id int64, removeData models.RemoveUserData) (error, []string) {
	err := db.AuthRemove(id, PasswordHash(removeData.Password))
	if err != models.IncorrectDataError {
		return models.IncorrectDataError, []string{"password"}
	}
	return err, nil
}
