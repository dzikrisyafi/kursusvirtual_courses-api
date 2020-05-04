package courses_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *sql.DB

	username = os.Getenv("mysql_username")
	password = os.Getenv("mysql_password")
	host     = os.Getenv("mysql_host")
	schema2  = os.Getenv("mysql_schema2")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema2,
	)
	var err error
	client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}

func DbConn() *sql.DB {
	return client
}
