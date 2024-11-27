package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"

	utils "gps_backend/utils"
	middleware "gps_backend/middleware"
	storage "gps_backend/storage"
	handlers "gps_backend/handlers"
)

func main() {
	storage.InitRedis()
	
    router := mux.NewRouter()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("UserID")
		token, err := utils.GenerateJWT(userID)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(token))
	}).Methods("GET")


	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.CheckAuthToken)
	protected.HandleFunc("/devices", handlers.GetDevices).Methods("GET")
	protected.HandleFunc("/preferences", handlers.GetPreferences).Methods("GET")
	protected.HandleFunc("/preferences", handlers.UpdatePreferences).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}