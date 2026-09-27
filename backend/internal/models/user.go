package models

import "time"
type User struct {
 ID int64 `json:"id"`
 Username string `json:"username"`
 QuotaGB int64 `json:"quota_gb"`
 ExpiresAt time.Time `json:"expires_at"`
}
