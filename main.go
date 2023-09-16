package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
type Vendor struct {
	ID               int    `json:"id"`
	CanteenID        int    `json:"canteen_id"`
	Name             string `json:"name"`
	OwnerID          sql.NullBool    `json:"owner_id"`
	OpeningTimestamp string `json:"opening_timestamp"`
	ClosingTimestamp string `json:"closing_timestamp"`
	Status           Status `json:"status"`
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
	return canteens
}

func getVendors() []*Vendor {
	db, err := sql.Open("mysql", "menu:root@tcp(db:3306)/menu_services_db")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to MySQL")
	results, err := db.Query("SELECT * FROM vendors")
	if err != nil {
		panic(err.Error())
	}

	var vendors []*Vendor
	for results.Next() {
		var u Vendor
		err = results.Scan(&u.ID, &u.CanteenID, &u.Name, &u.OwnerID, &u.OpeningTimestamp, &u.ClosingTimestamp, &u.Status)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(&u)
		vendors = append(vendors, &u)
	}
	fmt.Println(vendors)
	return vendors
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func canteenPage(w http.ResponseWriter, r *http.Request) {
	canteens := getCanteens()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(canteens)
}

func vendorPage(w http.ResponseWriter, r *http.Request) {
	vendors := getVendors()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(vendors)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/db/canteens", canteenPage)
	http.HandleFunc("/db/vendors", vendorPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
