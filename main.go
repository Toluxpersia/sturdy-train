package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// w.Write([]byte("Hello World"))
// }
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		switch r.URL.Path{
		case "/":
			w.Write([]byte("index page"))
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			w.Write([]byte("404 PAGE"))
			return
		}
	}
}

func main() {
	s := &api{addr: ":8080"}
	//s := &api{addr: ":8080"}
	//if err := http.Serve(s.addr, s); err != nil{
	//	log.Fatal(err)
	//}

	//s := &api{addr: ":8080"}
	//boot up and start server
	//dd

mux:= http.NewServeMux()

	// srv := &http.Server{
	// 	Addr: s.addr,
	// 	Handler: s,
	// }

	srv := &http.Server{
		Addr: s.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
