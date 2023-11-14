/*
 * omu-rest
 *
 * Laravel製RESTful APIサーバー。著者: [connect0459](https://github.com/connect0459)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"time"
)



// BookStock - BookStock Model
type BookStock struct {

	// ID
	Id int64 `json:"id,omitempty"`

	// 外部キー制約（types_branchesテーブル）
	TypeBranchId int64 `json:"type_branch_id,omitempty"`

	// 外部キー制約（books_infoテーブル）
	BookInfoId int64 `json:"book_info_id,omitempty"`

	// 未販売の在庫数
	Stock int64 `json:"stock,omitempty"`

	// 未処理の予約数
	Order int64 `json:"order,omitempty"`

	// 販売済みの在庫数
	Sold int64 `json:"sold,omitempty"`

	// レコード作成日
	CreatedAt time.Time `json:"created_at,omitempty"`

	// レコード更新日
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// AssertBookStockRequired checks if the required fields are not zero-ed
func AssertBookStockRequired(obj BookStock) error {
	return nil
}

// AssertBookStockConstraints checks if the values respects the defined constraints
func AssertBookStockConstraints(obj BookStock) error {
	return nil
}
