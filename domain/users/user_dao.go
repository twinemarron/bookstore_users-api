// dao: Data Access Object
// DB のアクセス処理を記述する
package users

import (
	"fmt"
	"strings"

	"github.com/twinemarron/bookstore_users-api/datasources/mysql/users_db"
	// "github.com/twinemarron/bookstore_users-api/domain/users"
	"github.com/twinemarron/bookstore_users-api/logger"
	"github.com/twinemarron/bookstore_users-api/utils/errors"
	"github.com/twinemarron/bookstore_users-api/utils/mysql_utils"
	// "github.com/twinemarron/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare get user statement", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to get user by id", getErr)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare save user statement", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to save user by id", saveErr)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()

	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to get last insert id after creating a user", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare update user statement", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to update user", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare delete user statement", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to delete user", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare find users by status statement", err)
		// クライアントへ返す error message
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to find users by status", err)
		// クライアントへ返す error message
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {

			// エラー内容をログとして記録
			logger.Error("error when scan user row into user struct", err)
			// クライアントへ返す error message
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		// エラー内容をログとして記録
		logger.Error("error when trying to prepare get user by email and password statement", err)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credencials")
		}

		// エラー内容をログとして記録
		logger.Error("error when trying to get user by email and password", getErr)
		// クライアントへ返す error message
		return errors.NewInternalServerError("database error")
	}
	return nil
}
