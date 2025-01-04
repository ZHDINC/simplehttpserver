package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReceivedMessage struct {
	SimpleString string
}

var global string

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
	mux.HandleFunc("/postMessage", postMessage)
	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Incoming Connection detected from %s", r.Host)
	w.Header().Set("content-type", "text/html")

}

func messageRoutine(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Incoming connection...")
	//fmt.Fprintln(w, "Connection successful!")
	fmt.Fprintln(w, global)
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var message ReceivedMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	fmt.Println(message)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	global += message.SimpleString + "\n"
	fmt.Println(message)
	w.WriteHeader(http.StatusOK)
}
