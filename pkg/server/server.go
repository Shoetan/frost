package server

import (
	"log"
	"net/http"

	"github.com/frost/pkg/server/handlers"
)


type APISERVER struct {
	addr string
}


func NewAPISERVER(addr string) *APISERVER  {

	return &APISERVER{
		addr: addr,
	}
	
}


func (s *APISERVER) Run() error {

	router := http.NewServeMux()

	router.HandleFunc("POST /generate", handlers.Request())

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Server is running on port %s", s.addr)

	return server.ListenAndServe()

} 

