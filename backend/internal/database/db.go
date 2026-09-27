package database

import (
 "context"
 "time"
 "github.com/jackc/pgx/v5/pgxpool"
)
func Open(url string) (*pgxpool.Pool,error) {
 cfg,err:=pgxpool.ParseConfig(url); if err!=nil{return nil,err}
 cfg.MaxConns=10; cfg.MinConns=1; cfg.MaxConnLifetime=time.Hour
 db,err:=pgxpool.NewWithConfig(context.Background(),cfg); if err!=nil{return nil,err}
 if err=db.Ping(context.Background()); err!=nil { db.Close(); return nil,err }
 _,err=db.Exec(context.Background(),`CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, username TEXT UNIQUE NOT NULL, quota_gb BIGINT NOT NULL, expires_at DATE NOT NULL, created_at TIMESTAMPTZ NOT NULL DEFAULT now())`)
 if err!=nil {db.Close();return nil,err}
 return db,nil
}
