package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/weareenvoy/nuface-recommendation-server-golang/config"
)

// Database instance
var DB *sql.DB

// Connect function
func Connect() error {
	var err error
	p := config.Config("DB_PORT")

	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	// Open db connections
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME")))

	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connection Opened to Database")
	return nil
}
