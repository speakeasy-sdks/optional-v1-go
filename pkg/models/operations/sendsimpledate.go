// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"net/http"
)

type SendSimpleDateRequest struct {
	Field      string            `queryParam:"style=form,explode=true,name=field"`
	SetToNull  bool              `queryParam:"style=form,explode=true,name=setToNull"`
	Simpledate shared.Simpledate `request:"mediaType=application/json"`
	UnSet      bool              `queryParam:"style=form,explode=true,name=unSet"`
}

func (o *SendSimpleDateRequest) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *SendSimpleDateRequest) GetSetToNull() bool {
	if o == nil {
		return false
	}
	return o.SetToNull
}

func (o *SendSimpleDateRequest) GetSimpledate() shared.Simpledate {
	if o == nil {
		return shared.Simpledate{}
	}
	return o.Simpledate
}

func (o *SendSimpleDateRequest) GetUnSet() bool {
	if o == nil {
		return false
	}
	return o.UnSet
}

type SendSimpleDateResponse struct {
	// HTTP response content type for this operation
	ContentType    string
	ServerResponse *shared.ServerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SendSimpleDateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendSimpleDateResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendSimpleDateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendSimpleDateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
