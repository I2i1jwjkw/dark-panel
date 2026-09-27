package main

import (
 "log"
 "net/http"
 "os"
 "dark-panel/backend/internal/config"
 "dark-panel/backend/internal/database"
 "dark-panel/backend/internal/handlers"
)

func main() {
 cfg := config.Load()
 db, err := database.Open(cfg.DatabaseURL)
 if err != nil { log.Fatal(err) }
 defer db.Close()
 mux := http.NewServeMux()
 mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type","application/json"); w.Write([]byte(`{"ok":true}`)) })
 mux.HandleFunc("POST /api/auth/login", handlers.Login(cfg))
 mux.HandleFunc("GET /api/users", handlers.Users(db))
 mux.HandleFunc("POST /api/users", handlers.CreateUser(db))
 mux.HandleFunc("POST /api/subscriptions", handlers.CreateSubscription(cfg))
 addr:=cfg.Addr
 if addr=="" { addr=":8080" }
 log.Printf("Dark Panel API listening on %s (env=%s)",addr,os.Getenv("APP_ENV"))
 log.Fatal(http.ListenAndServe(addr,mux))
}
