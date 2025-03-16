package utils

import (
	"encoding/json"
	"fmt"
	"gororoba/internal/controller"
	"log/slog"
	"net/http"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request) controller.HttpResponse

func HandleRequest(fn HttpHandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fn(w, r)

		if response.Headers != nil {
			for key, value := range response.Headers {
				w.Header().Set(key, value)
			}
		}

		if response.Headers["Content-Type"] == "" {
			w.Header().Set("Content-Type", "application/json")
		}

		if response.Body != nil {
			bytes, err := json.Marshal(response.Body)

			if err != nil {
				slog.Error(
					fmt.Sprintf("Error marshalling response: %v", err.Error()),
					slog.String("error", err.Error()),
				)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write(bytes)

			return
		}

		if response.StatusCode != 0 {
			w.WriteHeader(response.StatusCode)
		}
	}
}
