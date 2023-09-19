package controllers

import (
	"dineflow-menu-services/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllVendors(w http.ResponseWriter, r *http.Request) {
	results, err := models.GetAllVendors()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Vendor not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetVendorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID := vars["id"]

	vendor, err := models.GetVendorByID(vendorID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendor)
}

func CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendorRequest models.Vendor
	err := json.NewDecoder(r.Body).Decode(&vendorRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateVendor(vendorRequest)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Vendor created successfully")
}

func DeleteVendorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID := vars["id"]

	err := models.DeleteVendorByID(vendorID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Vendor deleted successfully")
}

func UpdateVendorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID := vars["id"]

	var updatedVendor models.Vendor
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedVendor); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := models.UpdateVendorByID(vendorID, updatedVendor)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Vendor updated successfully")
}