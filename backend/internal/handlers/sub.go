package handlers

import (
 "crypto/rand"
 "encoding/hex"
 "encoding/json"
 "net/http"
 "dark-panel/backend/internal/config"
)
func CreateSubscription(cfg config.Config) http.HandlerFunc {
 return func(w http.ResponseWriter,r *http.Request) {
  var in struct{Username string `json:"username"`}
  if json.NewDecoder(r.Body).Decode(&in)!=nil || in.Username=="" {http.Error(w,"invalid request",400);return}
  // Opaque identifier only; actual Xray provisioning is intentionally not implied.
  b:=make([]byte,24);if _,err:=rand.Read(b);err!=nil{http.Error(w,"random generation failed",500);return}
  url:=cfg.PublicURL+"/sub/"+hex.EncodeToString(b)
  w.Header().Set("Content-Type","application/json");w.WriteHeader(501)
  json.NewEncoder(w).Encode(map[string]string{"message":"subscription delivery is not configured","example_url":url})
 }
}
