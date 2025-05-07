package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Welcome endpoint
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": "Welcome"})
    }).Methods("GET")

    // Health endpoint
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
    }).Methods("GET")

    // Goodbye endpoint
    r.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
    }).Methods("GET")

    http.Handle("/", r)

    log.Println("Server starting on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}