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

type ErrorGeneratingRefreshToken struct {
	TypeError error
}

func (e ErrorGeneratingRefreshToken) Error() string {
	return fmt.Sprintf("Error generating refresh token: %w", e.TypeError)
}