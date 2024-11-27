package handler

import (
	"encoding/json"
	"net/http"
)

type HealthStatus string

const (
	UP   HealthStatus = "UP"
	DOWN HealthStatus = "DOWN"
)

type HealthCheckResponse struct {
	Status HealthStatus `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := HealthCheckResponse{Status: UP}
	json.NewEncoder(w).Encode(health)
}

func HealthCheckComplete(w http.ResponseWriter, r *http.Request) {
	health := HealthCheckResponse{Status: UP}
	json.NewEncoder(w).Encode(health)
}
