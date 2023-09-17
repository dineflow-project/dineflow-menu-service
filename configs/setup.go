package configs

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func ConnectDB() {
	Db, err = sql.Open("mysql", "menu:root@tcp(db:3306)/menu_services_db")
	if err != nil {
		log.Print(err.Error())
	}
	// defer Db.Close()

	fmt.Println("Connected to MySQL")
	// return Db
}
