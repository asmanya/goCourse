package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Serialization --> converting a go object into JSON string --> Marshal --> encoder
// Deserialization --> converting a JSON string into go object --> Unmarshal --> decoder

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	user := User{Name:"John", Email:"JohnSnow@example.com"}
	fmt.Println("Struct User format:", user)

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json string User format:", string(jsonData))

	var user1 User
	json.Unmarshal(jsonData, &user1)

	fmt.Println("Struct unmarshal User format:", user1)

	jsonData1 := `{"name": "Jeffery", "email": "Jeffery@example.com"}`
	reader := strings.NewReader(jsonData1)

	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Using decoder to read json string:", user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Using encoder to print go ds in json:", buf.String())
}
