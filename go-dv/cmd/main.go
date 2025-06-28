package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"dv/mvc/controllers"
	"dv/pkg/htemplate"
	"dv/pkg/server"
)

var (
	Port      int    = 9000
	Host      string = "0.0.0.0"
	hTemplate *htemplate.HTemplate
	hSrv *server.HttpServer
)

func main() {
	fmt.Printf("Server listening on %s:%d\n", Host, Port)

	viewsPath, err := os.Getwd()
	views := path.Join(viewsPath, "go-dv/mvc/views")
	fmt.Println("Views at: ", views)

	hTemplate, err := htemplate.NewHTemplate(views, "*.html")
	if err != nil {
		log.Fatalf("Unable to parse template files\n%s\n", err.Error())
	}

	hSrv = server.Start(Host, Port, "errors.log")

	// register routes
	hSrv.Register("/", controllers.Index(hTemplate))

	err = hSrv.ListenAndServe()
	if err != nil {
		hSrv.Info(err, "could not start server", true)
	}
}
