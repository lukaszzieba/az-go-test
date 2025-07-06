package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	res := map[string]any{"success": true}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
