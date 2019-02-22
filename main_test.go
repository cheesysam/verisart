package main

import (
	"bytes"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCertificateHandlerPost(t *testing.T) {

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

func TestCertificateHandlerPostBadJSON(t *testing.T) {
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

func TestDeleteCertificateDoesntExist(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/certificates/1", nil)
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

func TestDeleteCertificateDoesExist(t *testing.T) {
	loadElement()
	req, err := http.NewRequest("DELETE", "/certificates/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Certificate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func loadElement() {
	testData := `{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "", "Year": 1990, "Note": ""}`

	req, err := http.NewRequest("POST", "/certificates/1", bytes.NewBuffer([]byte(testData)))
	if err != nil {
		errors.New("load element error")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Certificate)

	handler.ServeHTTP(rr, req)

}
