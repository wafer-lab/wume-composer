package models

// 100-199 - success answer
// 200-299 - incorrect answer because of the user
// 300-399 - incorrect answer because developers

//easyjson:json
type MessageAnswer struct {
	Status  int    `json:"status, int" example:"100"`
	Message string `json:"message, string" example:"ok"`
}

//easyjson:json
type IncorrectDataAnswer struct {
	Status  int      `json:"status, int" example:"204"`
	Message string   `json:"message, string" example:"incorrect fields"`
	Data    []string `json:"data" example:"[email, username]"`
}

//easyjson:json
type UserDataAnswer struct {
	Status  int      `json:"status, int" example:"105"`
	Message string   `json:"message, string" example:"user found"`
	Data    UserData `json:"data"`
}

/* SUCCESS ANSWERS */

func GetSuccessAnswer(message string) *MessageAnswer {
	return &MessageAnswer{
		Status:  100,
		Message: message,
	}
}

var SignedInAnswer = MessageAnswer{
	Status:  101,
	Message: "signed in",
}

var SignedOutAnswer = MessageAnswer{
	Status:  102,
	Message: "signed out",
}

var SignedUpAnswer = MessageAnswer{
	Status:  103,
	Message: "signed up",
}

var PasswordUpdatedAnswer = MessageAnswer{
	Status:  104,
	Message: "password updated",
}

func GetUserDataAnswer(data UserData) *UserDataAnswer {
	return &UserDataAnswer{
		Status:  105,
		Message: "user found",
		Data:    data,
	}
}

var UserUpdatedAnswer = MessageAnswer{
	Status:  106,
	Message: "user updated",
}

var UserRemovedAnswer = MessageAnswer{
	Status:  107,
	Message: "user removed",
}

/* USERS ERRORS */

func GetUserErrorAnswer(error string) *MessageAnswer {
	return &MessageAnswer{
		Status:  200,
		Message: error,
	}
}

func GetIncorrectFieldsAnswer(data []string) *IncorrectDataAnswer {
	return &IncorrectDataAnswer{
		Status:  201,
		Message: "incorrect data",
		Data:    data,
	}
}

var NotSignedInAnswer = MessageAnswer{
	Status:  202,
	Message: "need be signed in",
}

var NotSignedOutAnswer = MessageAnswer{
	Status:  203,
	Message: "need be signed out",
}

var UserNotFoundAnswer = MessageAnswer{
	Status:  204,
	Message: "user not found",
}

/* DEVELOPERS ERRORS */

func GetDeveloperErrorAnswer(error string) *MessageAnswer {
	return &MessageAnswer{
		Status:  300,
		Message: error,
	}
}

var IncorrectJsonAnswer = MessageAnswer{
	Status:  301,
	Message: "incorrect JSON",
}
