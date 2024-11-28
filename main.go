package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"

	middleware "gps_backend/middleware"
	storage "gps_backend/storage"
	controllers "gps_backend/controllers"
)

func main() {
	storage.InitRedis()	
    router := mux.NewRouter()

	// Routes
	router.HandleFunc("/login", controllers.Login).Methods("GET")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.CheckAuthToken)
	protected.HandleFunc("/devices", controllers.GetDevices).Methods("GET")
	protected.HandleFunc("/preferences", controllers.GetPreferences).Methods("GET")
	protected.HandleFunc("/preferences", controllers.UpdatePreferences).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}