package main

import (
	"fmt"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/login", login)

	port := 8000
	log.Printf("Server run at localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}
