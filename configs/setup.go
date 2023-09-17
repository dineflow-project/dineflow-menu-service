package configs

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	dsn := os.Getenv("MYSQL_SOURCE")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}

	fmt.Println("Connected to MySQL")
}
