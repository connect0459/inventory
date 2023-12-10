// model/BookInfo.go
package model

import "time"

// @swagger:response BookInfo
type BookInfo struct {
	// ID
	// in: body
	// required: true
	ID uint64 `gorm:"type:bigint unsigned;primary_key" json:"id" example:"1"`
	// ISBN
	// in: body
	// required: true
	ISBN uint64 `gorm:"type:bigint unsigned not null" json:"isbn" example:"9784908434266"`
	// 書籍のタイトル
	// in: body
	// required: true
	Title string `gorm:"type:varchar(255) not null" json:"title" example:"これだけは知っておこう!情報リテラシー"`
	// 書籍の著者
	// in: body
	// required: true
	Author string `gorm:"type:varchar(255) not null" json:"author" example:"noa出版"`
	// 書籍の出版社
	// in: body
	// required: true
	Publisher string `gorm:"type:varchar(255) not null" json:"publisher" example:"noa出版"`
	// 書籍の説明
	// in: body
	// required: false
	Description string `gorm:"type:text" json:"description" example:"ハードウェア、ソフトウェア、ネットワーク、セキュリティ、情報モラルまで知っておきたい基礎知識が1冊に分かりやすくまとまっているテキストです。"`
	// ジャンルID
	// in: body
	// required: true
	TypeGenreID uint64 `gorm:"type:bigint unsigned not null" json:"genre" example:"1"`
	// 書籍の定価
	// in: body
	// required: true
	ListPrice uint64 `gorm:"type:int unsigned not null" json:"list_price" example:"1000"`
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
func (b *BookInfo) TableName() string {
	return "books_info"
}
