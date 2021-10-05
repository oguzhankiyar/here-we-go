package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	hub := NewHub()
	go hub.Run()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ws", hub.SocketHandler)

	fmt.Println("Listening http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}