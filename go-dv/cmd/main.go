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
	addr := fmt.Sprintf("%s:%d", Host, Port)

	fmt.Printf("Server listening on %s\n", addr)

	http.ListenAndServe(addr, nil)
}
