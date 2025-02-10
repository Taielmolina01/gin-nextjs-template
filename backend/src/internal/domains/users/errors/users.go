package errors

import (
	"fmt"
)

type ErrorUserNotExist struct {
	Email string
}

func (e ErrorUserNotExist) Error() string {
	return fmt.Sprintf("User with email %s is not registered", e.Email)
}

type ErrorUserAlreadyExist struct {
	Email string
}

func (e ErrorUserAlreadyExist) Error() string {
	return fmt.Sprintf("User with email %s is already registered", e.Email)
}

type ErrorUserMustHaveEmail struct{}

func (e ErrorUserMustHaveEmail) Error() string {
	return "User must have an email"
}

type ErrorUserMustHaveName struct{}

func (e ErrorUserMustHaveName) Error() string {
	return "User must have a name"
}

type ErrorPasswordMustHaveLenght8 struct{}

func (e ErrorPasswordMustHaveLenght8) Error() string {
	return "Password must have at least eight characters"
}

type ErrorUserRoleInvalid struct {
	Role string
}

func (e ErrorUserRoleInvalid) Error() string {
	return fmt.Sprintf("User role %s is not in the possible roles. \nPossible roles: 'user' and 'admin'", e.Role)
}

type ErrorWrongOldPassword struct{}

func (e ErrorWrongOldPassword) Error() string {
	return "The entered password is not the user's password"
}

type ErrorEncriptyngPassword struct{}

func (e ErrorEncriptyngPassword) Error() string {
	return "The entered password can't be encrypted"
}
