package handler

import (
    "github.com/dbiagi/gororoba/src/config"
    "log/slog"
)

type HealthCheckHandler struct {
    db config.Database
}

const (
    HealthStatusUP   = "UP"
    HealthStatusDOWN = "DOWN"
)

type HealthCheckResponse struct {
    Status   string         `json:"status"`
    Web      WebStatus      `json:"web"`
    Database DatabaseStatus `json:"database"`
}

type WebStatus struct {
    Status string `json:"status"`
}

type DatabaseStatus struct {
    Status string `json:"status"`
}

func NewHealthCheckHandler(db config.Database) HealthCheckHandler {
    return HealthCheckHandler{db: db}
}

func (h *HealthCheckHandler) Check() HealthCheckResponse {
    return HealthCheckResponse{Status: HealthStatusUP, Web: WebStatus{Status: HealthStatusUP}}
}

func (h *HealthCheckHandler) CheckComplete() HealthCheckResponse {
    err := h.db.Ping()
    if err != nil {
        slog.Error("Error pinging database: %v\n", err)
        return HealthCheckResponse{Status: HealthStatusDOWN, Web: WebStatus{Status: HealthStatusUP}, Database: DatabaseStatus{Status: HealthStatusDOWN}}
    }

    return HealthCheckResponse{Status: HealthStatusUP, Web: WebStatus{Status: HealthStatusUP}, Database: DatabaseStatus{Status: HealthStatusUP}}
}
