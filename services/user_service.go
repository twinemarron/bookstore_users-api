package services

import (
	"github.com/twinemarron/bookstore_users-api/domain/users"
	"github.com/twinemarron/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
	// var defaultUser users.User
	// return &defaultUser, nil
}
