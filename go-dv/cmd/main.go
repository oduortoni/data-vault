package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"dv/pkg/config"
	"dv/internal/users"
	"dv/mvc/controllers"
	"dv/mvc/models"
	"dv/pkg/auth"
	"dv/pkg/htemplate"
	"dv/pkg/server"
)

func main() {
	// Load application configuration
	cfg := config.New()
	fmt.Printf("Server listening on %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	// in-memory user model
	// userRepository := models.NewInternalUserRepository()

	// gorm user model
	db, err := gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		// Use log.Fatalf for cleaner exit on startup errors
		log.Fatalf("failed to connect database: %v", err)
	}
	userRepository := models.NewGormUserRepository(db)

	// create user service
	userService := users.NewUserService(userRepository)
	auth := auth.NewAuth(userService, []byte(cfg.Auth.JWTSecret), cfg.Auth.TokenCookieName)

	// fetch all neded directories
	cwd, err := os.Getwd()
	viewsDir := path.Join(cwd, "go-dv/mvc/views")
	staticDir := path.Join(cwd, "go-dv/www/static/")

	hTemplate, err := htemplate.NewHTemplate(viewsDir, "*.html")
	if err != nil {
		log.Fatalf("Unable to parse template files\n%s\n", err.Error())
	}

	hSrv := server.Start(cfg.Server.Host, cfg.Server.Port, "errors.log")

	// register routes
	hSrv.Register("GET", "/", controllers.Index(hTemplate))
	hSrv.Register("GET", "/static/", controllers.Static(staticDir))
	hSrv.Register("POST", "/auth/register", auth.Register)
	hSrv.Register("POST", "/auth/login", auth.Login)
	hSrv.Register("POST", "/auth/logout", auth.Logout)
	hSrv.Register("POST", "/auth/refresh", auth.Refresh)
	hSrv.Register("GET", "/dashboard", auth.AuthMiddleware(controllers.Dashboard))

	err = hSrv.ListenAndServe()
	if err != nil {
		hSrv.Info(err, "could not start server", true)
	}
}
