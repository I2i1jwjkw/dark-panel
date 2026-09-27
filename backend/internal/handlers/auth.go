package handlers

import (
 "encoding/json"
 "net/http"
 "dark-panel/backend/internal/config"
 "github.com/golang-jwt/jwt/v5"
 "golang.org/x/crypto/bcrypt"
 "time"
)
func Login(cfg config.Config) http.HandlerFunc {
 hash,err:=bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword),bcrypt.DefaultCost)
 if err!=nil { hash=[]byte{} }
 return func(w http.ResponseWriter,r *http.Request) {
  var in struct{ Username,Password string }; if json.NewDecoder(r.Body).Decode(&in)!=nil {http.Error(w,"invalid request",400);return}
  if cfg.AdminUsername=="" || cfg.JWTSecret=="" || bcrypt.CompareHashAndPassword(hash,[]byte(in.Password))!=nil || in.Username!=cfg.AdminUsername {http.Error(w,"unauthorized",401);return}
  token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{"sub":in.Username,"exp":time.Now().Add(12*time.Hour).Unix()})
  signed,e:=token.SignedString([]byte(cfg.JWTSecret));if e!=nil{http.Error(w,"token error",500);return}
  w.Header().Set("Content-Type","application/json");json.NewEncoder(w).Encode(map[string]string{"token":signed})
 }
}
