package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/1", handleRoute1)

	r.Post("/2", handleRoute2)

	r.Delete("/3", handleRoute3)

	fmt.Println("Starting server on port :3000")
	http.ListenAndServe(":3000", r)
}
func handleRoute1(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	_, _ = w.Write([]byte("Обработка маршрута 1\n"))
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"time":   time.Since(start),
	}).Info("Handled request")
}
func handleRoute2(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	_, _ = w.Write([]byte("Обработка маршрута 2\n"))
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"time":   time.Since(start),
	}).Info("Handled request")
}
func handleRoute3(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	_, _ = w.Write([]byte("Обработка маршрута 3\n"))
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"time":   time.Since(start),
	}).Info("Handled request")
}
