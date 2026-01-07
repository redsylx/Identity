package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	randomHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		number := rand.Intn(100)
		json.NewEncoder(w).Encode(map[string]int{"number": number})
	}

	http.HandleFunc("/random", enableCORS(randomHandler))

	println("Server running on http://0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
