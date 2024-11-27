package handlers

import (
	"encoding/json"
	"net/http"
	
	utils "gps_backend/utils"
	models "gps_backend/models"
)

var apiKey = "4wEeK_l4KkSmK9Oil3KxrKJfI_ZqCGKmVhVBMnUFD30"

func GetDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := utils.FetchDevices(apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch devices", http.StatusInternalServerError)
		return
	}

	var response []models.Device
	for _, device := range devices {
		latestPoint := device["latest_device_point"].(map[string]interface{})
		response = append(response, models.Device{
			DeviceID:   device["device_id"].(string),
			DisplayName: device["display_name"].(string),
			ActiveState: device["active_state"].(string),
			Lat:         latestPoint["lat"].(float64),
			Lng:         latestPoint["lng"].(float64),
			Speed:       latestPoint["speed"].(float64),
			DriveStatus: latestPoint["device_state"].(map[string]interface{})["drive_status"].(string),
		})
	}

	json.NewEncoder(w).Encode(response)
}
