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
	r.HandleFunc("/canteens", controllers.CreateCanteen).Methods("POST")
	r.HandleFunc("/vendors", controllers.CreateVendor).Methods("POST")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.UpdateCanteenByID).Methods("PUT", "PATCH")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.UpdateVendorByID).Methods("PUT", "PATCH")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.DeleteCanteenByID).Methods("DELETE")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.DeleteVendorByID).Methods("DELETE")
	http.Handle("/", r)
}
