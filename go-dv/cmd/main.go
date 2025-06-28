package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	Port int    = 9000
	Host string = "0.0.0.0"
)

func main() {
	fmt.Printf("Server listening on %s:%d\n", Host, Port)

	server, err := StartServer(Host, Port, "errors.log")
	if err != nil {
		server.Info(err, "could not start server", true)
	}
}

type Server struct {
	Host     string
	Port     int
	FileName string
	Errors   []string
}

func StartServer(host string, port int, filename string) (*Server, error) {
	server := &Server{
		Host:     host,
		Port:     port,
		FileName: filename,
	}

	server.Info(nil, "some error", true)

	err := server.ListenAndServe()

	return server, err
}

func (s Server) ListenAndServe() (err error) {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	err = http.ListenAndServe(addr, nil)
	return
}

func (s Server) Info(err error, msg string, save bool) {
	logMsg := fmt.Sprintf("[INFO] %s: %v\n", msg, err)
	fmt.Print(logMsg)

	if save {
		s.appendToFile(logMsg)
	}
}

func (s *Server) Panic(err error, msg string, save bool) {
	logMsg := fmt.Sprintf("[PANIC] %s: %v\n", msg, err)
	fmt.Fprint(os.Stderr, logMsg)

	if save {
		s.appendToFile(logMsg)
	}

	os.Exit(1)
}

func (s *Server) appendToFile(logMsg string) {
	if s.FileName == "" {
		return
	}

	f, err := os.OpenFile(s.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(logMsg); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}
