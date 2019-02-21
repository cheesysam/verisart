package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func Certificate(w http.ResponseWriter, r *http.Request) {
	testcert := certificate{ID: "1", Title: "i am a title", CreatedAt: 0, OwnerID: "", Year: 1990, Note: ""}
	log.Println("Cert handler")

	out, err := json.Marshal(testcert)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(out))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/certificate", Certificate)

	log.Fatal(http.ListenAndServe(":8000", router))
}
