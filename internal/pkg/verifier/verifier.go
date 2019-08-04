package verifier

import (
	"regexp"
)

var (
	emailRegexp    = regexp.MustCompile(`^[\w\-.]+@[\w\-.]+\.[a-zA-Z]{2,6}$`)
	usernameRegexp = regexp.MustCompile(`^[\w.]+$`)
)

func IsEmpty(value string) bool {
	return value == ""
}

func IsEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func IsEmailOrEmpty(email string) bool {
	return IsEmpty(email) || emailRegexp.MatchString(email)
}

func IsUsername(username string) bool {
	return usernameRegexp.MatchString(username)
}

func IsUsernameOrEmpty(username string) bool {
	return IsEmpty(username) || usernameRegexp.MatchString(username)
}
