package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Starting App")

	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Start server
	http.ListenAndServe(":9100", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API")
}
