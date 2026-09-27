package handlers

import (
 "context"
 "encoding/json"
 "net/http"
 "time"
 "github.com/jackc/pgx/v5/pgxpool"
 "dark-panel/backend/internal/models"
)
func Users(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter,r *http.Request) {
  rows,err:=db.Query(context.Background(),"SELECT id,username,quota_gb,expires_at FROM users ORDER BY id DESC")
  if err!=nil {http.Error(w,"database error",500);return};defer rows.Close()
  out:=[]models.User{}
  for rows.Next(){var u models.User;if rows.Scan(&u.ID,&u.Username,&u.QuotaGB,&u.ExpiresAt)!=nil{http.Error(w,"database error",500);return};out=append(out,u)}
  w.Header().Set("Content-Type","application/json");json.NewEncoder(w).Encode(out)
 }
}
func CreateUser(db *pgxpool.Pool) http.HandlerFunc {
 return func(w http.ResponseWriter,r *http.Request) {
  var in struct{Username string `json:"username"`;QuotaGB int64 `json:"quota_gb"`;ExpiresAt string `json:"expires_at"`}
  if json.NewDecoder(r.Body).Decode(&in)!=nil || in.Username=="" || len(in.Username)>64 || in.QuotaGB<1 || in.QuotaGB>100000 {http.Error(w,"invalid user",400);return}
  expiry,err:=time.Parse("2006-01-02",in.ExpiresAt);if err!=nil{http.Error(w,"invalid expiry",400);return}
  var u models.User
  err=db.QueryRow(context.Background(),"INSERT INTO users(username,quota_gb,expires_at) VALUES($1,$2,$3) RETURNING id,username,quota_gb,expires_at",in.Username,in.QuotaGB,expiry).Scan(&u.ID,&u.Username,&u.QuotaGB,&u.ExpiresAt)
  if err!=nil{http.Error(w,"could not create user",409);return}
  w.Header().Set("Content-Type","application/json");w.WriteHeader(201);json.NewEncoder(w).Encode(u)
 }
}
