package main

import (
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"

	"dineflow-menu-services/configs"
	"dineflow-menu-services/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()
	router := mux.NewRouter()
	routes.ProtectedRoute(router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
