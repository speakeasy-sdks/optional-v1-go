// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/optional-v1-go/pkg/types"
)

type Simpledate struct {
	Date         *types.Date `json:"date,omitempty"`
	DateNullable *types.Date `json:"dateNullable,omitempty"`
}

func (o *Simpledate) GetDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *Simpledate) GetDateNullable() *types.Date {
	if o == nil {
		return nil
	}
	return o.DateNullable
}