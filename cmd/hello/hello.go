package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	fmt.Println("Running on port", port)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
