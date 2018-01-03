package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name", "foo"}`))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name", "bar"}`))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	l, _ := net.Listen("tcp", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/foo/", foo)
	mux.HandleFunc("/bar/", bar)
	fmt.Println(http.Serve(l, mux))
}
