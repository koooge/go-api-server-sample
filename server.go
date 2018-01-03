package main

import (
	"fmt"
	"net"
	"net/http"
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
	l, _ := net.Listen("tcp", ":8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/foo/", foo)
	mux.HandleFunc("/bar/", bar)
	fmt.Println(http.Serve(l, mux))
}
