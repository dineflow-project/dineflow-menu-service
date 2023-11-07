package controllers

import (
	"dineflow-menu-services/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllMenus(w http.ResponseWriter, r *http.Request) {
	canteen := r.URL.Query().Get("canteen")
	vendor := r.URL.Query().Get("vendor")
	minprice, _ := strconv.ParseFloat(r.URL.Query().Get("minprice"), 64)
	maxprice, _ := strconv.ParseFloat(r.URL.Query().Get("maxprice"), 64)
	fmt.Println(canteen, vendor, minprice, maxprice)

	results, err := models.GetAllMenus(canteen, vendor, minprice, maxprice)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Menu not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetAllMenusByVendorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID := vars["id"]

	menus, err := models.GetAllMenusByVendorID(vendorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Menu not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func GetMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]

	menu, err := models.GetMenuByID(menuID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menuRequest models.Menu
	err := json.NewDecoder(r.Body).Decode(&menuRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = models.CreateMenu(menuRequest)
	if err != nil {
		log.Print(err.Error())

		// Handle the specific error returned from the model
		w.WriteHeader(http.StatusBadRequest)
		if models.IsVendorNotFoundError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Menu created successfully")
}

func DeleteMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]

	err := models.DeleteMenuByID(menuID)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Menu deleted successfully")
}

func UpdateMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]

	var updatedMenu models.Menu
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedMenu); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := models.UpdateMenuByID(menuID, updatedMenu)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Menu updated successfully")
}
