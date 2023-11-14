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



// TypePayment - TypePayment Model
type TypePayment struct {

	// ID
	Id int64 `json:"id,omitempty"`

	// 支払い方法
	Type string `json:"type,omitempty"`

	// レコード作成日
	CreatedAt time.Time `json:"created_at,omitempty"`

	// レコード更新日
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// AssertTypePaymentRequired checks if the required fields are not zero-ed
func AssertTypePaymentRequired(obj TypePayment) error {
	return nil
}

// AssertTypePaymentConstraints checks if the values respects the defined constraints
func AssertTypePaymentConstraints(obj TypePayment) error {
	return nil
}