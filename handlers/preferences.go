package handlers

import (
	"encoding/json"
	"net/http"
	storage "gps_backend/storage"
	models "gps_backend/models"
)

func GetPreferencesHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	if userID == "" {
		http.Error(w, "User-ID header is missing", http.StatusBadRequest)
		return
	}

	preferences, err := storage.GetPreferences(userID)
	if err != nil {
		http.Error(w, "Failed to get preferences", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(preferences)
}

func UpdatePreferencesHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID") 
	if userID == "" {
		http.Error(w, "User-ID header is missing", http.StatusBadRequest)
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

	w.WriteHeader(http.StatusOK)
}