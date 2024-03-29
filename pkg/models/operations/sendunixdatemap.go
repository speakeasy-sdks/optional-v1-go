// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/optional-v1-go/pkg/models/shared"
	"net/http"
)

type SendUnixDateMapRequest struct {
	Field       string             `queryParam:"style=form,explode=true,name=field"`
	SetToNull   bool               `queryParam:"style=form,explode=true,name=setToNull"`
	UnSet       bool               `queryParam:"style=form,explode=true,name=unSet"`
	Unixdatemap shared.Unixdatemap `request:"mediaType=application/json"`
}

func (o *SendUnixDateMapRequest) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *SendUnixDateMapRequest) GetSetToNull() bool {
	if o == nil {
		return false
	}
	return o.SetToNull
}

func (o *SendUnixDateMapRequest) GetUnSet() bool {
	if o == nil {
		return false
	}
	return o.UnSet
}

func (o *SendUnixDateMapRequest) GetUnixdatemap() shared.Unixdatemap {
	if o == nil {
		return shared.Unixdatemap{}
	}
	return o.Unixdatemap
}

type SendUnixDateMapResponse struct {
	// HTTP response content type for this operation
	ContentType    string
	ServerResponse *shared.ServerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SendUnixDateMapResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendUnixDateMapResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendUnixDateMapResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendUnixDateMapResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
