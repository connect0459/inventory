package model

import (
	"time"
)

// @swagger:response BookStock
type BookStock struct {
	// ID
	// in: body
	// required: true
	// example: 1
	ID uint64 `gorm:"type:bigint unsigned;primary_key" json:"id" example:"1"`
	// TypeBranchの支部ID
	// in: body
	// required: true
	// example: 1
	TypeBranchID uint64 `gorm:"type:int unsigned not null" json:"type_branch_id" example:"1"`
	// BookInfoの書籍ID
	// in: body
	// required: true
	// example: 1
	BookInfoID uint64 `gorm:"type:int unsigned not null" json:"book_info_id" example:"1"`
	// 書籍の支部ごとの販売価格
	// in: body
	// required: false
	// example: 500
	SalePrice uint64 `gorm:"type:int unsigned" json:"sale_price" example:"500"`
	// 在庫総数
	// in: body
	// required: true
	// example: 4
	Stock int `gorm:"type:int not null" json:"stock" example:"4"`
	// 予約総数
	// in: body
	// required: true
	// example: 1
	Order int `gorm:"type:int not null" json:"order" example:"1"`
	// 販売総数
	// in: body
	// required: true
	// example: 1
	Sold int `gorm:"type:int not null" json:"sold" example:"1"`
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
func (b *BookStock) TableName() string {
	return "books_stocks"
}
