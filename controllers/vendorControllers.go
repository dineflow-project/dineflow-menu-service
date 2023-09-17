package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"
	"dineflow-menu-services/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Status string

const (
	OPEN  Status = "Open"
	CLOSE Status = "Close"
)

type Vendor struct {
	ID               int          `json:"id"`
	CanteenID        int          `json:"canteen_id"`
	Name             string       `json:"name"`
	OwnerID          sql.NullBool `json:"owner_id"`
	OpeningTimestamp string       `json:"opening_timestamp"`
	ClosingTimestamp string       `json:"closing_timestamp"`
	Status           Status       `json:"status"`
}

func getAllVendors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected to MySQL")
	results, err := configs.Db.Query("SELECT * FROM vendors")
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
	err := configs.Db.QueryRow("SELECT id, canteen_id, name, owner_id, opening_timestamp, closing_timestamp, status FROM vendors WHERE ID = ?", vendorID).Scan(&vendor.ID, &vendor.CanteenID, &vendor.Name, &vendor.OwnerID, &vendor.OpeningTimestamp, &vendor.ClosingTimestamp, &vendor.Status)
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
	_, err = configs.Db.Exec("INSERT INTO vendors (canteen_id, name, owner_id, opening_timestamp, closing_timestamp, status) VALUES (?, ?, ?, ?, ?, ?)",
		vendorRequest.CanteenID, vendorRequest.Name, vendorRequest.OwnerID, vendorRequest.OpeningTimestamp, vendorRequest.ClosingTimestamp, vendorRequest.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Vendor created successfully")
}

func deleteVendorById(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	vendorID := vars["id"]

	// Perform the DELETE operation
	_, err := configs.Db.Exec("DELETE FROM vendors WHERE ID = ?", vendorID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Vendor deleted successfully")

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
	_, err := configs.Db.Exec(
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
