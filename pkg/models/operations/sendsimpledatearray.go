// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"net/http"
)

type SendSimpleDateArrayRequest struct {
	Field           string                 `queryParam:"style=form,explode=true,name=field"`
	SetToNull       bool                   `queryParam:"style=form,explode=true,name=setToNull"`
	Simpledatearray shared.Simpledatearray `request:"mediaType=application/json"`
	UnSet           bool                   `queryParam:"style=form,explode=true,name=unSet"`
}

func (o *SendSimpleDateArrayRequest) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *SendSimpleDateArrayRequest) GetSetToNull() bool {
	if o == nil {
		return false
	}
	return o.SetToNull
}

func (o *SendSimpleDateArrayRequest) GetSimpledatearray() shared.Simpledatearray {
	if o == nil {
		return shared.Simpledatearray{}
	}
	return o.Simpledatearray
}

func (o *SendSimpleDateArrayRequest) GetUnSet() bool {
	if o == nil {
		return false
	}
	return o.UnSet
}

type SendSimpleDateArrayResponse struct {
	ContentType    string
	ServerResponse *shared.ServerResponse
	StatusCode     int
	RawResponse    *http.Response
}

func (o *SendSimpleDateArrayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendSimpleDateArrayResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendSimpleDateArrayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendSimpleDateArrayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}