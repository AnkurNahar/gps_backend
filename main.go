package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

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

	// CORS Middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your frontend's origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "UserID"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}