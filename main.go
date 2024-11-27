package main

import( 
	"log"
	"net/http"
	"github.com/gorilla/mux"
	handlers "gps_backend/handlers"
)

func main() {
    router := mux.NewRouter()
	router.HandleFunc("/devices", handlers.GetDevicesHandler).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}