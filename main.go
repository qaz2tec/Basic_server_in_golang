package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static "))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", handleform)
	http.HandleFunc("/hello", handlehello)

	fmt.Printf("server is started at 8080 port\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleform(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Parseform() err: %v", err)
	}
	fmt.Fprint(w, "Post request sucessfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "name is %s\n", name)
	fmt.Fprint(w, "address is %s\n", address)
}

func handlehello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported ", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello!")
}
