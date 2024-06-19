package main

import "net/http"

func main() {
	// статический подход ограничен
	//fileDir := "course2/practice/rest_API_5/static"
	//http.ListenAndServe(":8080", http.FileServer(http.Dir(fileDir)))

	fileDir := "course2/practice/rest_API_5/static"
	fs := http.FileServer(http.Dir(fileDir))
	http.Handle("/", fs)
}
