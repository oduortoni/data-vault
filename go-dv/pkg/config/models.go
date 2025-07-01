package config

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type DatabaseConfig struct {
	DSN string
}

type AuthConfig struct {
	JWTSecret       string
	TokenCookieName string
}
