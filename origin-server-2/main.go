package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func origin(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[origin server] received request at: %s\n", time.Now())

	responseString := "Hello, from origin 2!"
	jsonResponse := map[string]string{"message": responseString}
	jsonData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response body
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", origin).Methods("GET")
	fmt.Println("origin 2 server running on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))

}
