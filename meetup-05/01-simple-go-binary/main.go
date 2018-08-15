package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>COS Gophers On The Web!!!</h1>")
	log.Println("page requested:", r.RequestURI)
}

func main() {
	log.Println("COS Gophers")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		log.Println(err)
	}
}
