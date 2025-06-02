package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server startup failed...,err:%v\n", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello world"))
}
