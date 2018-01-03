package main

import (
	"fmt"
	"net"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bar")
}

func main() {
	l, _ := net.Listen("tcp", ":8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", foo)
	mux.HandleFunc("/bar", bar)
	fmt.Println(http.Serve(l, mux))
}
