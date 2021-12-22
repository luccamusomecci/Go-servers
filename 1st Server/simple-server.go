package main

import (
	"net/http"
)

func main() {

	//Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/portfolio", portHandler)

	//The server starts
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func portHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my portfolio"))
}
