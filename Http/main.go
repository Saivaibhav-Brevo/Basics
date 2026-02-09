package main

import "net/http"

// type server struct {
// 	address string
// }

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		w.Write([]byte("Hello World"))

// 	case "/about":
// 		w.Write([]byte("This is the about page"))

// 	default:
// 		w.Write([]byte("404 Not Found"))
// 	}

// }

type api struct {
	address string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

func main() {
	// s := &server{address: ":8080"}
	// http.ListenAndServe(s.address, s)

	api := &api{address: ":8080"}

	//Initiialise Mux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.address,
		Handler: mux,
	}

	//Register Handlers
	mux.HandleFunc("GET /user/get", api.getUserHandler)
	mux.HandleFunc("POST /user/create", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
