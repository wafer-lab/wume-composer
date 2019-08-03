package auth

import (
	"net/http"
	"time"

	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/models"
)

func SetAuthCookie(w http.ResponseWriter, data models.JwtData) error {
	lifetime := time.Duration(config.Auth.CookieLifetime) * time.Minute

	jwtStr, err := data.Marshal(lifetime, []byte(config.Auth.Secret))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     config.Auth.CookieName,
		Value:    jwtStr,
		Expires:  time.Now().Add(lifetime),
		HttpOnly: true,
	})
	return nil
}

func RemoveAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     config.Auth.CookieName,
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
	})
}

func CheckJwt(token string) (models.JwtData, error) {
	data := models.JwtData{}
	err := data.UnMarshal(token, []byte(config.Auth.Secret))
	return data, err
}
