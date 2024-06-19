package main

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	// Применение middleware для логирования с помощью zap logger
	r.Use(LoggerMiddleware)

	r.Get("/1", handleRoute1)

	r.Post("/2", handleRoute2)

	r.Delete("/3", handleRoute3)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	logger, _ := zap.NewProduction()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("Handled request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("time", time.Since(start)),
		)
	})
}
func handleRoute1(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Обработка маршрута 1\n"))
}
func handleRoute2(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Обработка маршрута 2\n"))
}
func handleRoute3(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Обработка маршрута 3\n"))
}
