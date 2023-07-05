package web

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", helloHandler)
	log.Println("Started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
