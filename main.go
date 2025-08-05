package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Server starting on port 8080...")
	fmt.Println("Visit http://localhost:8080 to see Hello World!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
