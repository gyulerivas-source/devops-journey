package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"mensaje": "Hola desde binario compilado en Go", "nodo": "%s"}`, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor Go escuchando en :8080...")
	http.ListenAndServe(":8080", nil)
}
