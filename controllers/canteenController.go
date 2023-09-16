package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Status string

const (
	OPEN  Status = "Open"
	CLOSE Status = "Close"
)

type Canteen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getCanteens() []*Canteen {
	db, err := sql.Open("mysql", "menu:root@tcp(db:3306)/menu_services_db")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to MySQL")
	results, err := db.Query("SELECT * FROM canteens")
	if err != nil {
		panic(err.Error())
	}

	var canteens []*Canteen
	for results.Next() {
		var u Canteen
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(&u)
		canteens = append(canteens, &u)
	}
	fmt.Println(canteens)
	// c.IndentedJSON(http.StatusOK, canteens)
	return canteens
}