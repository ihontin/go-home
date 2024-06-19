package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"studentgit.kata.academy/Alkolex/go-kata/course2/practice/rest_API/internal/user"
	"time"
)

func main() {
	log.Println("Create router")
	router := httprouter.New()
	log.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	// адрес 127.0.0.1:1234 - это адрес лукбэк (lo) интерфейса, трафик возвращается, можно локально дебажить
	log.Println("Start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server is listening port :1234")
	log.Fatal(server.Serve(listener))
}
