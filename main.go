package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/profile/{name}", profile)

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Welcome to Basic GOLANG"}`))
}
func profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"My name is %s"}`, name)
}
