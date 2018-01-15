package users

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

var (
	errUnableToCreateUser                  = errors.New("unable to create user")
	errUnableToCreateUserDuplicateEmail    = errors.New("unable to create user: email already exists")
	errUnableToCreateUserDuplicateUsername = errors.New("unable to create user: username already exists")
	errUnableToProcessUserPayload          = errors.New("unable to process user payload")
	errUnknownProcessing                   = errors.New("unable process request")
	errUserNotFound                        = errors.New("user not found")
)

func errorStatusCode(err error) int {
	switch err {
	// 404
	case errUserNotFound:
		return http.StatusNotFound
	// 422
	case errUnableToCreateUserDuplicateEmail:
		fallthrough
	case errUnableToCreateUserDuplicateUsername:
		return http.StatusUnprocessableEntity
	}
	// 500
	return http.StatusInternalServerError
}

func storeErrors(err error) error {
	log.Println(err)
	if strings.Contains(err.Error(), "users_email_key") {
		return errUnableToCreateUserDuplicateEmail
	}
	if strings.Contains(err.Error(), "users_username_key") {
		return errUnableToCreateUserDuplicateUsername
	}
	return errUnableToCreateUser
}
