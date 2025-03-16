package utils

import (
	"gororoba/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	assertions func(w *httptest.ResponseRecorder)
	response   controller.HttpResponse
}

func TestHandleRequest(t *testing.T) {
	tests := []testCase{
		{
			name: "given a custom header, it should return the custom header",
			response: controller.HttpResponse{
				StatusCode: 200,
				Headers: map[string]string{
					"X-Custom-Header": "custom",
				},
			},
			assertions: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, "custom", w.Header().Get("X-Custom-Header"))
				assert.Equal(t, 200, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			},
		},
		{
			name: "given a status code, it should return that status code",
			response: controller.HttpResponse{
				StatusCode: 201,
			},
			assertions: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, 201, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			},
		},
		{
			name: "given a body, it should return a serialized json",
			response: controller.HttpResponse{
				StatusCode: 200,
				Body: map[string]string{
					"key": "value",
				},
			},
			assertions: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, `{"key":"value"}`, w.Body.String())
				assert.Equal(t, 200, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			},
		},
		{
			name: "given an error while marshalling the body, it should return a 500 status code",
			response: controller.HttpResponse{
				StatusCode: 200,
				Body: map[string]interface{}{
					"key": make(chan int),
				},
			},
			assertions: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, 500, w.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := HandleRequest(func(w http.ResponseWriter, r *http.Request) controller.HttpResponse {
				return tt.response
			})

			r := httptest.NewRecorder()
			fn(r, nil)
			tt.assertions(r)
		})
	}
}
