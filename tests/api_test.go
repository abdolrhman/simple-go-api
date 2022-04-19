package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/umangraval/Go-Mongodb-REST-boilerplate/routes"
)

var rou = routes.Routes()

func TestGetPeopleEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/trivia", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "success") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
