package errors

import (
	"fmt"
)

type ErrorUserTokenNotExist struct {
	UserEmail string
}

func (e ErrorUserTokenNotExist) Error() string {
	return fmt.Sprintf("User with email %s does not have a token", e.UserEmail)
}

type ErrorSigningToken struct {
	TypeError error
}

func (e ErrorSigningToken) Error() string {
	return fmt.Sprintf("Error signing token: %w", e.TypeError)
}

type ErrorGeneratingRefreshToken struct {}

func (e ErrorGeneratingRefreshToken) Error() string {
	return "Error generating the refresh token"
}