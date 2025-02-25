package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "myapp/Controllers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") 
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/news", Controllers.GetAllNews).Methods("GET")
    r.HandleFunc("/change", Controllers.UpdateNews).Methods("PUT")

    http.Handle("/", r)

    handler := enableCORS(r)

    fmt.Println("Сервер запущен на http://localhost:8080")
    var connect = http.ListenAndServe(":8080", handler)
    if connect != nil {
        log.Fatalf("Не удалось запустить сервер: %v", connect)
    }
}
