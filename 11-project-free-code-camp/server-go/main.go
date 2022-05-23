package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
	}

	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not support", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Staring server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
