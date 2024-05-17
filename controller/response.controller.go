package controller

import (
	"encoding/json"
	"net/http"
	"van_thailand_server/models"
)

func ReturnFailed(w http.ResponseWriter) {
	resp := models.Response{
		Message: "fail",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func ReturnSuccess(w http.ResponseWriter) {
	resp := models.Response{
		Message: "success",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
