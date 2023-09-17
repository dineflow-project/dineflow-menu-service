package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"dineflow-menu-services/configs"
	"dineflow-menu-services/models"
	"dineflow-menu-services/routes"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	configs.ConnectDB()
	if err := models.AutoMigrateDB(); err != nil {
		log.Fatal("Database migration error:", err)
	}

	router := mux.NewRouter()
	routes.ProtectedRoute(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
