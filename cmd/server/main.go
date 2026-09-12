package main

import (
 "encoding/json"
 "log"
 "net/http"
)

func main() {
 http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  _ = json.NewEncoder(w).Encode(map[string]string{"status":"ok","contractVersion":"2026-09"})
 })
 log.Fatal(http.ListenAndServe(":8080", nil))
}
