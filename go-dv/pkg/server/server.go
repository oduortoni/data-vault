package server

import (
	"fmt"
	"net/http"
	"os"
)

type HttpServer struct {
	Host     string
	Port     int
	FileName string
	Errors   []string
	Router   *Router
}

func Start(host string, port int, filename string) *HttpServer {
	router := NewRouter()

	server := &HttpServer{
		Host:     host,
		Port:     port,
		FileName: filename,
		Router:   router,
	}

	return server
}

func (s HttpServer) ListenAndServe() (err error) {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	err = http.ListenAndServe(addr, s.Router)
	return
}

func (s HttpServer) Register(method, pattern string, handler http.HandlerFunc) {
	s.Router.register(method, pattern, handler)
}

func (s HttpServer) Info(err error, msg string, save bool) {
	logMsg := fmt.Sprintf("[INFO] %s: %v\n", msg, err)
	fmt.Print(logMsg)

	if save {
		s.appendToFile(logMsg)
	}
}

func (s *HttpServer) Panic(err error, msg string, save bool) {
	logMsg := fmt.Sprintf("[PANIC] %s: %v\n", msg, err)
	fmt.Fprint(os.Stderr, logMsg)

	if save {
		s.appendToFile(logMsg)
	}

	os.Exit(1)
}

func (s *HttpServer) appendToFile(logMsg string) {
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
