package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet - prints a greeting
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s!", name)
}

// MyGreeterHandler - handles incomming HTTP requests
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
