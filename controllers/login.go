package controllers

import (
	"encoding/json"
	"net/http"
	
	utils "gps_backend/utils"
	models "gps_backend/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("UserID")
	token, err := utils.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"authToken":  token,
	}

	response := models.Response{
		Status:  "success",
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}