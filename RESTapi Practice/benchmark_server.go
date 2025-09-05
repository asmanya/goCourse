package RESTapiPractice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Person Struct
type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Sample Data
var personData = map[string]Person {
	"1": {Name: "John Snow", Age: 25},
	"2": {Name: "Jeffery", Age: 15},
	"3": {Name: "Sansa", Age: 45},
}

// Handler function for the endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	// Check if the ID exists in the personData map
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the person data to JSON and write to the response
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func benchmarkServer() {
	// Define the port
	port := 8080

	// Print the confirmation message
	fmt.Println("Server started on port:", port)

	// Set up the endpoint and the handler function
	http.HandleFunc("/person", getPersonHandler)

	// Start the server
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// using htop to monitor the threads
// command used --> wrk -t8 -c400 -d30s "http://localhost:8080/person?id=1"