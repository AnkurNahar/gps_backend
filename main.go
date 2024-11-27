package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"
	storage "gps_backend/storage"
	handlers "gps_backend/handlers"
)

func main() {
	storage.InitRedis()
	
    router := mux.NewRouter()
	router.HandleFunc("/devices", handlers.GetDevicesHandler).Methods("GET")
	router.HandleFunc("/preferences", handlers.GetPreferencesHandler).Methods("GET")
	router.HandleFunc("/preferences", handlers.UpdatePreferencesHandler).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}