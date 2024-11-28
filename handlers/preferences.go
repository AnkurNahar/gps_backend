package handlers

import (
	"encoding/json"
	"net/http"
	
	storage "gps_backend/storage"
	models "gps_backend/models"
	middleware "gps_backend/middleware"
) 

func GetPreferences(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	if userID == "" {
		http.Error(w, "UserID missing", http.StatusBadRequest)
		return
	}

	preferences, err := storage.GetPreferences(userID)
	if err != nil {
		http.Error(w, "Failed to get preferences", http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Status: "success",
		Data: preferences,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdatePreferences(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	if userID == "" {
		http.Error(w, "UserID missing", http.StatusBadRequest)
		return
	}

	var preferences models.Preferences
	if err := json.NewDecoder(r.Body).Decode(&preferences); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := storage.SavePreferences(userID, preferences)
	if err != nil {
		http.Error(w, "Failed to update preferences", http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Status: "success",
	}

	json.NewEncoder(w).Encode(response)
}