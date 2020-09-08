package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestControllerIndex(t *testing.T) {
	tests := []struct {
		name     string
		route    string
		method   string
		expected int
	}{
		{
			name:     "Get API Info",
			route:    "/",
			method:   "GET",
			expected: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			router := httprouter.New()
			router.GET(tt.route, IndexControlller)
			router.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.expected {
				t.Errorf("%v, returned wrong status code: got %v want %v",
					tt.name, status, tt.expected)
			}
		})
	}
}
