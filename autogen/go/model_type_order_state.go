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



// TypeOrderState - TypeOrderState Model
type TypeOrderState struct {

	// ID
	Id int64 `json:"id,omitempty"`

	// 予約、販売済み、キャンセル、返品などのステータス
	Type string `json:"type,omitempty"`

	// レコード作成日
	CreatedAt time.Time `json:"created_at,omitempty"`

	// レコード更新日
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// AssertTypeOrderStateRequired checks if the required fields are not zero-ed
func AssertTypeOrderStateRequired(obj TypeOrderState) error {
	return nil
}

// AssertTypeOrderStateConstraints checks if the values respects the defined constraints
func AssertTypeOrderStateConstraints(obj TypeOrderState) error {
	return nil
}
