package main

import (
	"fmt"
	"log"
	"net/http"
)

func routeValidator(handler http.HandlerFunc, path string, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}
		if r.Method != method {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}
		handler(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	routeValidator(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello!")
	}, "/hello", "GET")(w, r)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	routeValidator(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	}, "/form", "POST")(w, r)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
