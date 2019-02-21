package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCertificateHandler(t *testing.T) {

	testData := `{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "", "Year": 1990, "Note": ""}`

	req, err := http.NewRequest("POST", "/certificates/1", bytes.NewBuffer([]byte(testData)))
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

	expected := `Cert Posted`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCertificateHandlerBadJSON(t *testing.T) {
	testData := `{"Bad Data": "1"}`

	req, err := http.NewRequest("POST", "/certificates/1", bytes.NewBuffer([]byte(testData)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Certificate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
