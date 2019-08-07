package models

import (
	"wume-composer/internal/pkg/common/verifier"
)

/********************
 *    IN MODELS     *
 ********************/

/* SIGN IN DATA */

//easyjson:json
type SignInData struct {
	Login    string `json:"login" example:"test@mail.ru"`
	Password string `json:"password" example:"Qwerty123"`
}

func (v SignInData) Validate() (incorrectFields []string) {
	if !verifier.IsEmail(v.Login) && !verifier.IsUsername(v.Login) {
		incorrectFields = append(incorrectFields, "login")
	}
	if verifier.IsEmpty(v.Password) {
		incorrectFields = append(incorrectFields, "password")
	}
	return
}

/* SIGN UP DATA */

//easyjson:json
type SignUpData struct {
	Email    string `json:"email" example:"user_test@test.com"`
	Username string `json:"username" example:"user_test"`
	Password string `json:"password" example:"SecretPass1!"`
}

func (v SignUpData) Validate() (incorrectFields []string) {
	if !verifier.IsEmail(v.Email) {
		incorrectFields = append(incorrectFields, "email")
	}
	if !verifier.IsUsername(v.Username) {
		incorrectFields = append(incorrectFields, "username")
	}
	if verifier.IsEmpty(v.Password) {
		incorrectFields = append(incorrectFields, "password")
	}
	return
}

/* UPDATE USER DATA */

//easyjson:json
type UpdateUserData struct {
	Email    string `json:"email" example:"user_test@test.com"`
	Username string `json:"username" example:"user_test"`
}

func (data UpdateUserData) Validate() (incorrectFields []string) {
	if !verifier.IsEmail(data.Email) {
		incorrectFields = append(incorrectFields, "email")
	}
	if !verifier.IsUsername(data.Username) {
		incorrectFields = append(incorrectFields, "username")
	}
	return
}

/* UPDATE PASSWORD DATA */

//easyjson:json
type UpdatePasswordData struct {
	OldPassword     string `json:"old_password" example:"SecretPass1!"`
	NewPassword     string `json:"new_password" example:"SecretPass2!"`
	PasswordConfirm string `json:"password_confirm" example:"SecretPass2!"`
}

func (v UpdatePasswordData) Validate() (incorrectFields []string) {
	if verifier.IsEmpty(v.OldPassword) {
		incorrectFields = append(incorrectFields, "old_password")
	}
	if verifier.IsEmpty(v.NewPassword) {
		incorrectFields = append(incorrectFields, "new_password")
	}
	if v.PasswordConfirm != v.NewPassword {
		incorrectFields = append(incorrectFields, "password_confirm")
	}
	return
}

/* REMOVE USER DATA */

//easyjson:json
type RemoveUserData struct {
	Password string `json:"password" example:"SecretPass1!"`
}

func (v RemoveUserData) Validate() (incorrectFields []string) {
	if verifier.IsEmpty(v.Password) {
		incorrectFields = append(incorrectFields, "password")
	}
	return
}

/********************
 *    OUT MODELS    *
 ********************/

/* USER DATA */

//easyjson:json
type UserData struct {
	Id       int64  `json:"id, string" example:"1"`
	Username string `json:"username, string" example:"user_test"`
	Email    string `json:"email, string" example:"user_test@test.com"`
	Avatar    string `json:"avatar, string" example:"/upload/avatar.jpg"`
}

//easyjson:json
type UsersRowData struct {
	Id       int64  `json:"id, string" example:"1"`
	Username string `json:"username, string" example:"user_test"`
	Email    string `json:"email, string" example:"user_test@test.com"`
	Avatar    string `json:"avatar, string" example:"/upload/avatar.jpg"`
}

//easyjson:json
type UsersData struct {
	Users []UsersRowData `json:"users"`
	Count uint64         `json:"count" example:"123"`
}
