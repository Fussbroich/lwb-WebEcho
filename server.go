package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	// ein neuer Multiplexer für das Routing
	mux := http.NewServeMux()

	// registriere einige Methoden und Routen, die Deine WebApp beantworten soll

	// GET /meinPfad
	mux.HandleFunc("GET /meinPfad",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handled GET Request an %s", r.URL.Path)
		})

	// GET an der Wurzel
	mux.HandleFunc("GET /",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handled GET Request an der Wurzel")
		})

	// POST an der Wurzel
	mux.HandleFunc("POST /",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handled POST Request an der Wurzel")
		})

	// Fallback Handler für alle anderen Requests
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Willkommen! Du hast den Server mit der %s Methode unter %s erreicht.", r.Method, r.URL.Path)
		})

	// Baue einen Web-Server und starte ihn
	server := &http.Server{
		Addr:    net.JoinHostPort("localhost", "8081"),
		Handler: http.Handler(mux),
	}
	fmt.Printf("starte Web-Server unter http://%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server-Fehler: %s\n", err)
	}
}
