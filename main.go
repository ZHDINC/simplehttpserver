package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":80",
		Handler: routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal()
	}
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))
	mux.Handle("/", http.StripPrefix("/", fs))
	//mux.HandleFunc("/", homePage)
	mux.HandleFunc("/messages", messageRoutine)
	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Incoming Connection detected from %s", r.Host)
	w.Header().Set("content-type", "text/html")

}

func messageRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming connection...")
	fmt.Fprintln(w, "Connection successful!")
}
