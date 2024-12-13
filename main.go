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
	preferenceStorage, err := storage.NewRedisStorage("localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}	
    router := mux.NewRouter()

	// Routes
	router.HandleFunc("/login", controllers.Login).Methods("GET")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.CheckAuthToken)
	protected.HandleFunc("/devices", controllers.GetDevices).Methods("GET")
	protected.HandleFunc("/preferences", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetPreferences(w, r, preferenceStorage)
	}).Methods("GET")
	protected.HandleFunc("/preferences", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdatePreferences(w, r, preferenceStorage)
	}).Methods("POST")

	// CORS Middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, 
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "UserID"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}