package main

import (
	"fmt"
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(fmt.Sprintf("{\"health\"}")))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", foo)
	mux.HandleFunc("/dog", bar)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
