package RESTapiPractice

import (
	"fmt"
	"io"
	"net/http"
)

func client() {

	// Create a new http client
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1") // mock api
	// another mock api "https://swapi.dev/api/people/1"

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// READ and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
	fmt.Println(body)

}