package models

// 100-199 - success answer
// 200-299 - incorrect answer because of the user
// 300-399 - incorrect answer because developers

type MessageAnswer struct {
	Status  int    `json:"status, int" example:"100"`
	Message string `json:"message, string" example:"ok"`
}

type IncorrectFieldsAnswer struct {
	Status  int      `json:"status, int" example:"204"`
	Message string   `json:"message, string" example:"incorrect fields"`
	Data    []string `json:"data" example:"[email, username]"`
}

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
	Status:  107,
	Message: "user updated",
}

var UserRemovedAnswer = MessageAnswer{
	Status:  108,
	Message: "user removed",
}

/* USERS ERRORS */

// For future use
//
// func GetUserErrorAnswer(error string) *MessageAnswer {
// 	return &MessageAnswer{
// 		Status:  200,
// 		ChatMessage: error,
// 	}
// }

var NotSignedInAnswer = MessageAnswer{
	Status:  201,
	Message: "not signed in",
}

var AlreadySignedInAnswer = MessageAnswer{
	Status:  202,
	Message: "already signed in",
}

var AlreadySignedOutAnswer = MessageAnswer{
	Status:  203,
	Message: "already signed out",
}

var UserNotFoundAnswer = MessageAnswer{
	Status:  205,
	Message: "user not found",
}

func GetIncorrectFieldsAnswer(data []string) *IncorrectFieldsAnswer {
	return &IncorrectFieldsAnswer{
		Status:  204,
		Message: "incorrect fields",
		Data:    data,
	}
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
