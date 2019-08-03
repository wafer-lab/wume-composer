package auth

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/models"
)

func CreateAuthCookie(data models.JwtData, lifetime time.Duration) *http.Cookie {
	jwtStr, err := data.Marshal(lifetime, []byte(config.Auth.Secret))
	if err != nil {
		return &http.Cookie{}
	}

	return &http.Cookie{
		Name:     config.Auth.CookieName,
		Value:    jwtStr,
		Expires:  time.Now().Add(lifetime),
		HttpOnly: true,
	}
}

func CheckJwt(token string) (models.JwtData, error) {
	data := models.JwtData{}
	err := data.UnMarshal(token, []byte(config.Auth.Secret))
	return data, err
}

func PasswordHash(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func SignUp(signUpData models.SignUpData) (models.JwtData, error, []string) {
	return models.JwtData{}, nil, nil
}

func SignIn(signInData models.SignInData) (models.JwtData, error, []string) {
	return models.JwtData{}, nil, nil
}

func UpdateAuth(id int64, userData models.UpdateUserData) (models.JwtData, error, []string) {
	return models.JwtData{}, nil, nil
}

func UpdatePassword(id int64, passwordData models.UpdatePasswordData) (error, []string) {
	return nil, nil
}

func RemoveAuth(id int64, removeData models.RemoveUserData) (error, []string) {
	return nil, nil
}
