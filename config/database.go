package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const mysqlDSN = "root:qweasdzxc1@tcp(127.0.0.1:3306)/userdata"

// ConnectDB establishes a connection to MySQL database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Println("Error connecting to MySQL:", err)
		return nil, err
	}
	return db, nil
}
