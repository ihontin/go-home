package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/1", helloWorld)
	router.Post("/1", helloWorld)
	router.Put("/1", helloWorld)

	router.Get("/2", helloWorld2)
	router.Post("/2", helloWorld2)
	router.Put("/2", helloWorld2)

	router.Get("/3", helloWorld3)
	router.Post("/3", helloWorld3)
	router.Put("/3", helloWorld3)

	router.HandleFunc("/", notFound404)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
func notFound404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("Not Found"))
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello world"))
}
func helloWorld2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello world 2"))
}
func helloWorld3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello world 3"))
}
