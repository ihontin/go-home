package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/url"
)

func main() {
	r := chi.NewRouter()

	r.Route("/group1", func(r chi.Router) {
		r.Get("/1", helloWorld)
		r.Get("/2", helloWorld2)
		r.Get("/3", helloWorld3)
	})

	r.Route("/group2", func(r chi.Router) {
		r.Get("/1", helloWorld)
		r.Get("/2", helloWorld2)
		r.Get("/3", helloWorld3)
	})

	r.Route("/group3", func(r chi.Router) {
		r.Get("/1", helloWorld)
		r.Get("/2", helloWorld2)
		r.Get("/3", helloWorld3)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	u, _ := url.Parse(r.URL.String())
	g := u.Path[1:]
	outString := []byte(fmt.Sprintf("Group %s Привет, мир 1\n", string(g[len(g)-3])))
	_, _ = w.Write(outString)
}
func helloWorld2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	u, _ := url.Parse(r.URL.String())
	g := u.Path[1:]
	outString := []byte(fmt.Sprintf("Group %s Привет, мир 2\n", string(g[len(g)-3])))
	_, _ = w.Write(outString)
}
func helloWorld3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	u, _ := url.Parse(r.URL.String())
	g := u.Path[1:]
	outString := []byte(fmt.Sprintf("Group %s Привет, мир 3\n", string(g[len(g)-3])))
	_, _ = w.Write(outString)
}
