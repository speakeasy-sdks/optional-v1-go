// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package optionalandnullable

import (
	"fmt"
	"github.com/speakeasy-sdks/optional-v1-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://localhost:3000",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient  HTTPClient
	SecurityClient HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// OptionalAndNullable
type OptionalAndNullable struct {
	Body       *body
	RfcDate    *rfcDate
	SimpleDate *simpleDate
	UnixDate   *unixDate

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*OptionalAndNullable)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *OptionalAndNullable) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *OptionalAndNullable) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *OptionalAndNullable) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *OptionalAndNullable) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *OptionalAndNullable {
	sdk := &OptionalAndNullable{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.1.0",
			GenVersion:        "2.86.10",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
	}

	sdk.Body = newBody(sdk.sdkConfiguration)

	sdk.RfcDate = newRfcDate(sdk.sdkConfiguration)

	sdk.SimpleDate = newSimpleDate(sdk.sdkConfiguration)

	sdk.UnixDate = newUnixDate(sdk.sdkConfiguration)

	return sdk
}