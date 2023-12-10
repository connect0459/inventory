// model/BookInfo.go
package model

import "time"

// @swagger:response TypeBranch
type TypeBranch struct {
	// 支部のID
	// in: body
	// required: true
	ID uint64 `gorm:"type:bigint unsigned;primary_key" json:"id" example:"1"`
	// 支部の名前
	// in: body
	// required: true
	Type string `gorm:"type:varchar(255) not null" json:"type" example:"○○大学"`
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
func (b *TypeBranch) TableName() string {
	return "types_branches"
}
