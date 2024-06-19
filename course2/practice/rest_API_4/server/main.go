package main

import (
	"fmt"
	"net/http"
)

// второй способ определения маршрутов
type httpHandler struct {
	message string
}

func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, h.message)
}

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "First page")
	})

	//выводе html файла
	fileDir := "course2/practice/rest_API_4/views/hello.html"
	http.HandleFunc("/file_html", func(resp http.ResponseWriter, req *http.Request) {
		http.ServeFile(resp, req, fileDir)
	})
	h2 := httpHandler{message: "Second page"}
	h3 := httpHandler{message: "third page"}

	http.Handle("/second", h2)
	http.Handle("/third", h3)

	http.ListenAndServe(":8080", nil)
}
