package main

import "net/http"

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Hello World"))

	case "/about":
		w.Write([]byte("This is the about page"))

	default:
		w.Write([]byte("404 Not Found"))
	}

}

func main() {
	s := &server{address: ":8080"}
	http.ListenAndServe(s.address, s)
}
