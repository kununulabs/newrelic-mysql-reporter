package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection returns a connection to our mysql database
func GetConnection(url, user, pass string) (*sql.DB, error) {
	if len(user) > 0 || len(pass) > 0 {
		url = user + ":" + pass + "@" + url
	}

	mysql, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}

	if err := mysql.Ping(); err != nil {
		return nil, err
	}

	return mysql, nil
}
