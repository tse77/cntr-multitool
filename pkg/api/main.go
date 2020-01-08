package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Run() {
	fmt.Println("Starting App")

	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/k8s-node", k8sNodeHandler).Methods("GET")

	// Start server
	http.ListenAndServe(":9100", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API")
}

// k8sNodeHandler ...
func k8sNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodename := os.Getenv("MY_NODE_NAME")
	podnamespace := os.Getenv("MY_POD_NAMESPACE")
	mypodip := os.Getenv("MY_POD_IP")

	fmt.Fprintf(w, "Nodename: %s\nPod namespace: %s\nPod IP: %s",
		nodename, podnamespace, mypodip)
}
