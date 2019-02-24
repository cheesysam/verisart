package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var CertificateDB []certificate

func Certificate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error: ", err)
		}
		owner := r.Header.Get("Authorization")
		PostCert(w, body, owner)
	}

	if r.Method == "DELETE" {
		DeleteCert(vars["id"], w)
	}

	if r.Method == "PATCH" {
		err := DeleteCert(vars["id"], w)
		if err != nil {
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error: ", err)
		}
		PostCert(w, body, "")
	}
}

func PostCert(w http.ResponseWriter, body []byte, owner string) error {
	//TODO check cert with id exists?

	fmt.Println(string(body))
	fmt.Println(owner) //TODO write owner to cert

	var newCert certificate
	err := json.Unmarshal(body, &newCert)
	if err != nil {
		fmt.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("bad request")
	}

	if newCert.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("bad request")
	}

	fmt.Println(CertificateDB)
	CertificateDB = append(CertificateDB, newCert)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Cert Posted")
	return nil
}

func DeleteCert(id string, w http.ResponseWriter) error {
	for i, cert := range CertificateDB {

		if cert.ID == id {
			CertificateDB = append(CertificateDB[:i], CertificateDB[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return nil
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return errors.New("bad request")

}

func UserCertificates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var matchingUserCertificates []certificate
	for _, cert := range CertificateDB {
		if cert.OwnerID == vars["userId"] {
			matchingUserCertificates = append(matchingUserCertificates, cert)
		}
	}
	out, err := json.Marshal(matchingUserCertificates)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))
	return
}

func CertificateTransfer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "CertificateTransferHandler")
}

func Reset(w http.ResponseWriter, r *http.Request) {
	CertificateDB = nil
	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/certificates/{id:[0-9a-zA-Z]+}", Certificate).Methods("POST", "DELETE", "PATCH")
	router.HandleFunc("/certificates/{id:[0-9a-zA-Z]+}/transfer", CertificateTransfer).Methods("PATCH")
	router.HandleFunc("/users/{userId:[0-9a-zA-Z]+}/certificates", UserCertificates).Methods("GET")
	router.HandleFunc("/reset", Reset) //Debug helper method in lieu of test database
	log.Fatal(http.ListenAndServe(":8000", router))
}
