package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dineflow-menu-services/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllCanteens(w http.ResponseWriter, r *http.Request) {
	results, err := models.GetAllCanteens()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Canteen not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetCanteenByID(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Query the database to get the canteen by ID using the new function
	canteen, err := models.GetCanteenByID(canteenID)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the canteen information to JSON and send it as the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(canteen)
}

func CreateCanteen(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var newCanteen models.Canteen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newCanteen); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := models.CreateCanteen(newCanteen)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Canteen created successfully")
}

func DeleteCanteenByID(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	err := models.DeleteCanteenByID(canteenID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Canteen deleted successfully")
}

func UpdateCanteenByID(w http.ResponseWriter, r *http.Request) {
	// Get the canteen ID from the URL path parameters
	vars := mux.Vars(r)
	canteenID := vars["id"]

	// Parse the incoming JSON request body into a Canteen struct
	var updatedCanteen models.Canteen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedCanteen); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Update the canteen in the database using the new function
	err := models.UpdateCanteenByID(canteenID, updatedCanteen)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Canteen updated successfully")
}
