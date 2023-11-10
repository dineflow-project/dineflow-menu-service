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
	canteenId := r.URL.Query().Get("canteenId")
	vendorId := r.URL.Query().Get("vendorId")
	if canteenId == "" {
		canteenId = "-1"
	}
	if vendorId == "" {
		vendorId = "-1"
	}
	canteenIdint, err := strconv.Atoi(canteenId)
	if err != nil {
		fmt.Println(err)
	}
	vendorIdint, err := strconv.Atoi(vendorId)
	if err != nil {
		fmt.Println(err)
	}
	minprice, _ := strconv.ParseFloat(r.URL.Query().Get("minprice"), 64)
	maxprice, _ := strconv.ParseFloat(r.URL.Query().Get("maxprice"), 64)
	// fmt.Println(canteenIdint, vendorIdint, minprice, maxprice)

	results, err := models.GetAllMenus(canteenIdint, vendorIdint, minprice, maxprice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Menu not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetAllMenusByVendorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vendorID := vars["id"]

	menus, err := models.GetAllMenusByVendorID(vendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Menu not found", http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menus)
}

func GetMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]

	menu, err := models.GetMenuByID(menuID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menuRequest models.Menu
	err := json.NewDecoder(r.Body).Decode(&menuRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	menu, err := models.CreateMenu(menuRequest)
	if err != nil {
		log.Print(err.Error())

		// Handle the specific error returned from the model
		w.WriteHeader(http.StatusBadRequest)
		if models.IsVendorNotFoundError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		// fmt.Fprintf(w, "Menu created successfully")
		json.NewEncoder(w).Encode(menu)
	}

}

func DeleteMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]

	err := models.DeleteMenuByID(menuID)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Menu updated successfully")
}
