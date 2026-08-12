package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {

	http.HandleFunc("/bubble/sequential", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == "OPTIONS" {
			return
		}

		size, _ := strconv.Atoi(r.URL.Query().Get("size"))
		if size == 0 {
			size = 50000
		}
		resp := RunBubbleSort(size, false)
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/bubble/parallel", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		if r.Method == "OPTIONS" {
			return
		}

		size, _ := strconv.Atoi(r.URL.Query().Get("size"))
		if size == 0 {
			size = 50000
		}
		resp := RunBubbleSort(size, true)
		json.NewEncoder(w).Encode(resp)
	})
	http.HandleFunc("/selection/sequential", func(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" { return }

	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if size == 0 { size = 50000 }
	resp := RunSelectionSort(size, false)
	json.NewEncoder(w).Encode(resp)
})

http.HandleFunc("/selection/parallel", func(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" { return }

	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if size == 0 { size = 50000 }
	resp := RunSelectionSort(size, true)
	json.NewEncoder(w).Encode(resp)
})

	fmt.Println(" API ready at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
