package mysql_utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/twinemarron/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewInternalServerError(fmt.Sprintf("no record matching giving id"), err)
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error parsing database response"), err)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
