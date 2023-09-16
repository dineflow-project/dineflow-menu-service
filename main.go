package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	ID               int          `json:"id"`
	CanteenID        int          `json:"canteen_id"`
	Name             string       `json:"name"`
	OwnerID          sql.NullBool `json:"owner_id"`
	OpeningTimestamp string       `json:"opening_timestamp"`
	ClosingTimestamp string       `json:"closing_timestamp"`
	Status           Status       `json:"status"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "menu:root@tcp(db:3306)/menu_services_db")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	// router := gin.Default()
	// router.GET("/canteen", getCanteens)
	// router.Run("localhost:8080")
	r := mux.NewRouter()
	r.HandleFunc("/canteens", getAllCanteens).Methods("GET")
	r.HandleFunc("/canteens/{id:[0-9]+}", getCanteenByID).Methods("GET")
	r.HandleFunc("/vendors", getAllVendors).Methods("GET")
	r.HandleFunc("/vendors/{id:[0-9]+}", getVendorByID).Methods("GET")
	r.HandleFunc("/canteens", createCanteen).Methods("POST")
	r.HandleFunc("/vendors", createVendor).Methods("POST")
	r.HandleFunc("/canteens/{id:[0-9]+}", updateCanteenById).Methods("PUT", "PATCH")
	r.HandleFunc("/vendors/{id:[0-9]+}", updateVendorById).Methods("PUT", "PATCH")
	r.HandleFunc("/canteens/{id:[0-9]+}", deleteCanteenById).Methods("DELETE")
	r.HandleFunc("/vendors/{id:[0-9]+}", deleteVendorById).Methods("DELETE")
	http.Handle("/", r)
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/db/canteens", canteenPage)
	// http.HandleFunc("/db/vendors", vendorPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAllCanteens(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all canteens
	rows, err := db.Query("SELECT ID, Name FROM canteens")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the canteens
	var canteens []Canteen

	// Iterate over the rows and scan the canteens
	for rows.Next() {
		var canteen Canteen
		err := rows.Scan(&canteen.ID, &canteen.Name)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		canteens = append(canteens, canteen)
	}

	// Serialize the canteens to JSON and send them as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(canteens)
}

func getCanteenByID(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Query the database to get the canteen by ID
	var canteen Canteen
	err := db.QueryRow("SELECT ID, Name FROM canteens WHERE ID = ?", canteenID).Scan(&canteen.ID, &canteen.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Canteen not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Serialize the canteen information to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(canteen)
}

func getAllVendors(w http.ResponseWriter, r *http.Request) {
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
	// Serialize the canteens to JSON and send them as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendors)
}

func getVendorByID(w http.ResponseWriter, r *http.Request) {
	// Get the vendor ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["id"]

	// Query the database to get the vendor by ID
	var vendor Vendor
	err := db.QueryRow("SELECT id, canteen_id, name, owner_id, opening_timestamp, closing_timestamp, status FROM vendors WHERE ID = ?", vendorID).Scan(&vendor.ID, &vendor.CanteenID, &vendor.Name, &vendor.OwnerID, &vendor.OpeningTimestamp, &vendor.ClosingTimestamp, &vendor.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Vendor not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Serialize the vendor information to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendor)
}

func createVendor(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data from the request body into a VendorRequest struct
	var vendorRequest Vendor
	err := json.NewDecoder(r.Body).Decode(&vendorRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the data into your database
	_, err = db.Exec("INSERT INTO vendors (canteen_id, name, owner_id, opening_timestamp, closing_timestamp, status) VALUES (?, ?, ?, ?, ?, ?)",
		vendorRequest.CanteenID, vendorRequest.Name, vendorRequest.OwnerID, vendorRequest.OpeningTimestamp, vendorRequest.ClosingTimestamp, vendorRequest.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Vendor created successfully")
}

func createCanteen(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var newCanteen Canteen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCanteen); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Println(newCanteen)
	fmt.Println(newCanteen.Name)

	// Perform the INSERT operation, assuming you have a table structure similar to your Canteen struct
	_, err := db.Exec("INSERT INTO canteens (Name) VALUES (?)", newCanteen.Name)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Canteen created successfully")
}

func deleteCanteenById(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Perform the DELETE operation
	_, err := db.Exec("DELETE FROM canteens WHERE ID = ?", canteenID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Canteen deleted successfully")
}

func deleteVendorById(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["id"]

	// Perform the DELETE operation
	_, err := db.Exec("DELETE FROM vendors WHERE ID = ?", vendorID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Vendor deleted successfully")

}

func updateCanteenById(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Parse the incoming JSON request body into a Canteen struct
	var updatedCanteen Canteen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedCanteen); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Perform the UPDATE operation
	_, err := db.Exec("UPDATE canteens SET Name = ? WHERE ID = ?", updatedCanteen.Name, canteenID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Canteen updated successfully")
}

func updateVendorById(w http.ResponseWriter, r *http.Request) {
	// Get the vendor ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["id"]

	// Parse the incoming JSON request body into a Vendor struct
	var updatedVendor Vendor
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedVendor); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Execute an UPDATE SQL statement to update the vendor's information
	_, err := db.Exec(
		"UPDATE vendors SET name = ?, owner_id = ?, opening_timestamp = ?, closing_timestamp = ?, status = ? WHERE id = ?",
		updatedVendor.Name, updatedVendor.OwnerID, updatedVendor.OpeningTimestamp, updatedVendor.ClosingTimestamp, updatedVendor.Status, vendorID,
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Vendor updated successfully")
}
