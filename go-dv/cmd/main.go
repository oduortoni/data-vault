package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"dv/internal/users"
	"dv/mvc/controllers"
	"dv/mvc/models"
	"dv/pkg/auth"
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

	userModel := models.NewUserModel()
	userService := users.NewUserService(userModel)
	auth := auth.NewAuth(userService, []byte("secret"), "access_token")

	// fetch all neded directories
	cwd, err := os.Getwd()
	viewsDir := path.Join(cwd, "go-dv/mvc/views")
	staticDir := path.Join(cwd, "go-dv/www/static/")

	hTemplate, err := htemplate.NewHTemplate(viewsDir, "*.html")
	if err != nil {
		log.Fatalf("Unable to parse template files\n%s\n", err.Error())
	}

	hSrv = server.Start(Host, Port, "errors.log")

	// register routes
	hSrv.Register("GET", "/", controllers.Index(hTemplate))
	hSrv.Register("GET", "/static/", controllers.Static(staticDir))
	hSrv.Register("POST", "/auth/register", auth.Register)
	hSrv.Register("POST", "/auth/login", auth.Login)
	hSrv.Register("POST", "/auth/logout", auth.Logout)
	hSrv.Register("POST", "/auth/refresh", auth.Refresh)
	hSrv.Register("GET", "/dashboard", auth.AuthMiddleware(controllers.Dashboard(hTemplate)))

	err = hSrv.ListenAndServe()
	if err != nil {
		hSrv.Info(err, "could not start server", true)
	}
}
