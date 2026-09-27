package config

import "os"

type Config struct { Addr, DatabaseURL, JWTSecret, AdminUsername, AdminPassword, PublicURL string }
func Load() Config {
 return Config{Addr:os.Getenv("APP_ADDR"),DatabaseURL:os.Getenv("DATABASE_URL"),JWTSecret:os.Getenv("JWT_SECRET"),AdminUsername:os.Getenv("ADMIN_USERNAME"),AdminPassword:os.Getenv("ADMIN_PASSWORD"),PublicURL:os.Getenv("PUBLIC_URL")}
}
