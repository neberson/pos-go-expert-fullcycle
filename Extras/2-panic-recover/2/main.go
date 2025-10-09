package main

import (
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r.(string))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}

}
