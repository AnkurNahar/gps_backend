package main

import( 
	"fmt"
	handlers "gps_backend/handlers"
)

func main() {
	response, err := handlers.GetDevicesHandler()

    fmt.Println(response)
}