package server

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		url := r.URL
		msg := url.Query().Get("id")
		fmt.Println(msg)
	}
}

func Serve() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for listening at :3000...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
