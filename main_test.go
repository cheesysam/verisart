package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TODO test for owner header set
func TestCertificateHandlerPost(t *testing.T) {

	testData := `{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "", "Year": 1990, "Note": ""}`

	req, err := http.NewRequest("POST", "/certificates/1", bytes.NewBuffer([]byte(testData)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CertificateHandler)

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
	handler := http.HandlerFunc(CertificateHandler)

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
	handler := http.HandlerFunc(CertificateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteCertificateDoesExist(t *testing.T) {
	loadElement("sam")
	req, err := http.NewRequest("DELETE", "/certificates/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CertificateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPatchCertificateDoesExist(t *testing.T) {
	loadElement("sam")
	testData := `{"ID": "2", "Title": "i am a patched title", "CreatedAt": 4, "OwnerID": "", "Year": 1890, "Note": "new note"}`
	req, err := http.NewRequest("PATCH", "/certificates/1", bytes.NewBuffer([]byte(testData)))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CertificateHandler)

	handler.ServeHTTP(rr, req)
}

func TestUserNoCerts(t *testing.T) {
	reset()
	req, err := http.NewRequest("GET", "/users/sam/certificates", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserCertificates)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != "null" {
		t.Errorf("body not correct: got %x expected %v", rr.Body.String(), "null")
	}
}

func TestUserWithCerts(t *testing.T) {
	loadElement("sam")
	req, err := http.NewRequest("GET", "/users/sam/certificates", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "sam"})
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserCertificates)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() == "null" {
		t.Errorf("body not correct: got %x expected some data", rr.Body.String())
	}
}

func loadElement(userID string) {
	testData := `{"ID": "1", "Title": "i am a title", "CreatedAt": 0, "OwnerID": "` + userID + `", "Year": 1990, "Note": ""}`

	req, err := http.NewRequest("POST", "/certificates/1", bytes.NewBuffer([]byte(testData)))
	req = mux.SetURLVars(req, map[string]string{"userId": "sam"})
	if err != nil {
		errors.New("load element error")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CertificateHandler)

	handler.ServeHTTP(rr, req)
}

func reset() {
	req, _ := http.NewRequest("GET", "/reset", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Reset)
	handler.ServeHTTP(rr, req)
}
