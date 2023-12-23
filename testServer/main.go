package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := map[string]string{
			"Read":     "/read",
			"ReadbyID": "/read/id",
			"Create":   "/create",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
	})

	log.Println("Starting Server...")
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Listening at port", port)

	http.Serve(listener, nil)
}
