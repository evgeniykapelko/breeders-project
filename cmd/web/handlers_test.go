package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/god-breeds", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Error("handler returned wrong status code:", status)
	}

}
