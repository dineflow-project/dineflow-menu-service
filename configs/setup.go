package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func ConnectDB() {
	Db, err = sql.Open(os.Getenv("MYSQL_DRIVER"), os.Getenv("MYSQL_SOURCE"))
	if err != nil {
		log.Print(err.Error())
	}
	// defer Db.Close()
	fmt.Println("Connected to MySQL")
	// return Db
}
