package routes

import (
	"net/http"

	// "github.com/gin-gonic/gin"

	"github.com/gorilla/mux"
)

func ProtectedRoute(r *mux.Router) {
	r.HandleFunc("/canteens", controllers.getAllCanteens).Methods("GET")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.getCanteenByID).Methods("GET")
	r.HandleFunc("/vendors", controllers.getAllVendors).Methods("GET")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.getVendorByID).Methods("GET")
	r.HandleFunc("/canteens", controllers.createCanteen).Methods("POST")
	r.HandleFunc("/vendors", controllers.createVendor).Methods("POST")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.updateCanteenById).Methods("PUT", "PATCH")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.updateVendorById).Methods("PUT", "PATCH")
	r.HandleFunc("/canteens/{id:[0-9]+}", controllers.deleteCanteenById).Methods("DELETE")
	r.HandleFunc("/vendors/{id:[0-9]+}", controllers.deleteVendorById).Methods("DELETE")
	http.Handle("/", r)
}
