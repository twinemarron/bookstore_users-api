package users_db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "MYSQL_USERS_USERNAME"
	mysql_users_password = "MYSQL_USERS_PASSWORD"
	mysql_users_host     = "MYSQL_USERS_HOST"
	mysql_users_schema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB

	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)
)

func getDataSourceName() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	username := os.Getenv(mysql_users_username)
	password := os.Getenv(mysql_users_password)
	host := os.Getenv(mysql_users_host)
	schema := os.Getenv(mysql_users_schema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	return dataSourceName
}

func init() {
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	username, password, host, schema,
	// )
	dataSourceName := getDataSourceName()
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successflly configured")
}
