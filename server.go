package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    // Handle GET requests
    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Received GET request on %s\n", r.URL)
    })

    // Handle POST requests
    mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading body", http.StatusInternalServerError)
            return
        }
        defer r.Body.Close()

        fmt.Fprintf(w, "Received POST request on %s\nBody:\n%s\n", r.URL, body)
    })

    // Start the server
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", mux)
}
