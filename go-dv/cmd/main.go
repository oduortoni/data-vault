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
	hSrv      *server.HttpServer
)

func main() {
	Port = server.Port(Port)
	fmt.Printf("Server listening on %s:%d\n", Host, Port)

	cwd, err := os.Getwd()
	viewsDir := path.Join(cwd, "go-dv/mvc/views")
	staticDir := path.Join(cwd, "go-dv/www/static/")
	fmt.Println("Static at", staticDir)

	hTemplate, err := htemplate.NewHTemplate(viewsDir, "*.html")
	if err != nil {
		log.Fatalf("Unable to parse template files\n%s\n", err.Error())
	}

	hSrv = server.Start(Host, Port, "errors.log")

	// register routes
	hSrv.Register("/", controllers.Index(hTemplate))
	hSrv.Register("/static/", controllers.Static(staticDir))

	err = hSrv.ListenAndServe()
	if err != nil {
		hSrv.Info(err, "could not start server", true)
	}
}
