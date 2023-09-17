package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dineflow-menu-services/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Canteen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCanteens(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all canteens
	rows, err := configs.Db.Query("SELECT ID, Name FROM canteens")
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

func GetCanteenByID(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Query the database to get the canteen by ID
	var canteen Canteen
	err := configs.Db.QueryRow("SELECT ID, Name FROM canteens WHERE ID = ?", canteenID).Scan(&canteen.ID, &canteen.Name)
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

func CreateCanteen(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var newCanteen Canteen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCanteen); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Perform the INSERT operation, assuming you have a table structure similar to your Canteen struct
	_, err := configs.Db.Exec("INSERT INTO canteens (Name) VALUES (?)", newCanteen.Name)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Canteen created successfully")
}

func DeleteCanteenById(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Perform the DELETE operation
	_, err := configs.Db.Exec("DELETE FROM canteens WHERE ID = ?", canteenID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Canteen deleted successfully")
}

func UpdateCanteenById(w http.ResponseWriter, r *http.Request) {
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
	_, err := configs.Db.Exec("UPDATE canteens SET Name = ? WHERE ID = ?", updatedCanteen.Name, canteenID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Canteen updated successfully")
}
