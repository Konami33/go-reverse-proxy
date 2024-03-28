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
	responseString := "Hello, from origin 1!"

	// Convert the string to a JSON object
	jsonResponse := map[string]string{"message": responseString}

	//Encoding of json data using marshal
	jsonData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response body
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {

	//creating a router
	r := mux.NewRouter()
	r.HandleFunc("/", origin).Methods("GET")
	fmt.Println("origin 1 server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
