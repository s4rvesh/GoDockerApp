package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Docker basic app")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Docker")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
