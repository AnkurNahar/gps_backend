package main

import( 
	"fmt"
	"gps-backend/handlers"
)

func main() {
	response, err := handlers.GetDevicesHandler()

    fmt.Println(response)
}