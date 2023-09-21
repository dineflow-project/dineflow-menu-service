package routes

import (
	"net/http"

	"dineflow-menu-services/controllers"

	"github.com/gorilla/mux"
)

func ProtectedRoute(r *mux.Router) {
	r.HandleFunc("/canteens", controllers.GetAllCanteens).Methods("GET")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.GetCanteenByID).Methods("GET")
	r.HandleFunc("/vendors", controllers.GetAllVendors).Methods("GET")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.GetVendorByID).Methods("GET")
	r.HandleFunc("/vendors/canteens/{id:[0-9]+}", controllers.GetAllVendorsByCanteenID).Methods("GET")
	r.HandleFunc("/canteens", controllers.CreateCanteen).Methods("POST")
	r.HandleFunc("/vendors", controllers.CreateVendor).Methods("POST")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.UpdateCanteenByID).Methods("PUT", "PATCH")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.UpdateVendorByID).Methods("PUT", "PATCH")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.DeleteCanteenByID).Methods("DELETE")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.DeleteVendorByID).Methods("DELETE")

	r.HandleFunc("/menus", controllers.GetAllMenus).Methods("GET")
	r.HandleFunc("/menus", controllers.CreateMenu).Methods("POST")
	r.HandleFunc("/menus/vendors/{id:[0-9]+}", controllers.GetAllMenusByVendorID).Methods("GET")
	r.HandleFunc("/menus/{id:[0-9]+}", controllers.DeleteMenuByID).Methods("DELETE")
	r.HandleFunc("/menus/{id:[0-9]+}", controllers.UpdateMenuByID).Methods("PUT", "PATCH")

	http.Handle("/", r)
}
