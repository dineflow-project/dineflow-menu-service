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

func GetAllMenus(w http.ResponseWriter, r *http.Request) {
	results, err := models.GetAllMenus()
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
        return
    }

    err = models.CreateMenu(menuRequest)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Menu created successfully")
}

func DeleteMenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["id"]
    fmt.Fprintln(menuID)

	err := models.DeleteMenuByID(menuID)
    fmt.Fprintln(err)
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

	err := models.UpdateMenuByID(MenuID, updatedMenu)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Menu updated successfully")
}
