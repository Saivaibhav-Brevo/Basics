package main

import "net/http"

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	s := &server{address: ":8080"}
	http.ListenAndServe(s.address, s)
}
