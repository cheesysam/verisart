package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var CertificateDB []certificate

func Certificate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//TODO check cert with id exists?

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("error: ", err)
		}

		fmt.Println(string(body))

		var newCert certificate
		err = json.Unmarshal(body, &newCert)
		if err != nil {
			fmt.Println("error: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCert.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(CertificateDB)
		CertificateDB = append(CertificateDB, newCert)

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Cert Posted")
	}
}

func UserCertificates(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "RetrievingUserCertificates")
	//out, err := json.Marshal(testcert)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
}

func CertificateTransfer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "CertificateTransferHandler")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/certificates/{id:[0-9a-zA-Z]+}", Certificate).Methods("POST", "DELETE", "PATCH")
	router.HandleFunc("/certificates/{id:[0-9a-zA-Z]+}/transfer", CertificateTransfer).Methods("PATCH")
	router.HandleFunc("/users/{userId:[0-9a-zA-Z]+}/certificates", UserCertificates).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
