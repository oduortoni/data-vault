package main

import (
	"fmt"
	"net/http"
)

var (
	Port int = 9000
	Host string = "0.0.0.0"
)

func main() {
	fmt.Printf("Server listening on %s:%d\n", Host, Port)
	
	server, err := StartServer(Host, Port, "")
	if err != nil {
		server.Info(err, "could not start server", false)
	}
}

type Server struct {
	Host string
	Port int
	FileName string
	Errors []string
}

func StartServer(host string, port int, filename string) (*Server, error) {
	server := &Server{
		Host: host,
		Port: port,
		FileName: filename,
	}

	err := server.ListenAndServe()

	return server, err
}

func (s Server) ListenAndServe() (err error) {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	err = http.ListenAndServe(addr, nil)
	return
}

func (s Server) Info(err error, msg string, save bool) {
	if save {
		//save to log file
	}
}

func (s *Server) Panic(err error, msg string, save bool) {
	if save {
		// save to file
	}
}
