package RESTapiPractice

import (
	"fmt"
	"log"
	"net/http"
)

func server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello server!")
	})

	const port string = ":3000"

	fmt.Println("Server is listening on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}