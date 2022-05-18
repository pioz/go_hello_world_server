package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "World"
	}
	log.Println("Get", name)
	fmt.Fprintf(w, "Hello %s!", name)
}
