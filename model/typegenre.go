// model/BookInfo.go
package model

import "time"

// @swagger:response TypeGenre
type TypeGenre struct {
	// ジャンルID
	// in: body
	// required: true
	ID uint64 `gorm:"type:bigint unsigned;primary_key" json:"id" example:"1"`
	// ジャンル名
	// in: body
	// required: true
	Type string `gorm:"type:varchar(255) not null" json:"type" example:"情報学"`
	// レコードの作成日時
	// in: body
	// required: false
	CreatedAt time.Time `json:"created_at" example:"2023-06-20T21:20:14.000000Z"`
	// レコードの更新日時
	// in: body
	// required: false
	UpdatedAt time.Time `json:"updated_at" example:"2023-06-20T21:20:14.000000Z"`
	// レコードの削除日時
	// in: body
	// required: false
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty" example:"null"`
}

// TableName sets the insert table name for this struct type
func (b *TypeGenre) TableName() string {
	return "types_genres"
}
