package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"dineflow-menu-services/configs"
	"dineflow-menu-services/models"
	"dineflow-menu-services/routes"

	"github.com/gorilla/handlers"
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
	if errMg := models.AutoMigrateDB(); errMg != nil {
		log.Fatal("Database migration error:", errMg)
	}

	// Create a CORS handler
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}), // Adjust the allowed origins as needed
	)
	router := mux.NewRouter()
	routes.ProtectedRoute(router)
	routerWithCORS := corsHandler(router)

	local_os := runtime.GOOS
	if local_os == "windows" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), routerWithCORS))
	} else {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), routerWithCORS))
	}
}
