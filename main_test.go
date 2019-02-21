package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCertificateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/certificate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Certificate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "", "Year": 1990, "Note": ""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
