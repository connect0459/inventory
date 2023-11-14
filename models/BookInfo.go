package model

import "time"

type BookInfo struct {
//ここ大文字にすることをお忘れなく
    Id              int              `gorm:"primary_key" json:"id"`
    First_name      string           `json:"first_name"`
    Family_name     string           `json:"family_name"`   
    Email           string           `json:"email"`
    Created_at      time.Time        `json:"created_at"`
    Updated_at      time.Time        `json:"updated_at"`
    Deleted_at      *time.Time       `json:"deleted_at"`
}
