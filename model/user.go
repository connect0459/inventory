package model

import "time"

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Area       string    `json:"area"`
	Prefecture string    `json:"prefecture"`
	Url        string    `json:"url"`
	BikeName   string    `json:"bike_name"`
	Experience int8      `json:"experience"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// Spots      []Spot    `json:"posts" gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" param:"user_id"`
}
// TableName sets the insert table name for this struct type
func (b *User) TableName() string {
	return "users"
}
