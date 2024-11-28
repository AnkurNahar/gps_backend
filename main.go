package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"

	middleware "gps_backend/middleware"
	storage "gps_backend/storage"
	handlers "gps_backend/handlers"
)

func main() {
	storage.InitRedis()
	
    router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login).Methods("GET")


	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.CheckAuthToken)
	protected.HandleFunc("/devices", handlers.GetDevices).Methods("GET")
	protected.HandleFunc("/preferences", handlers.GetPreferences).Methods("GET")
	protected.HandleFunc("/preferences", handlers.UpdatePreferences).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}