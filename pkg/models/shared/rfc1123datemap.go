// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Rfc1123datemap struct {
	DateTime  map[string]string `json:"dateTime,omitempty"`
	DateTime1 map[string]string `json:"dateTime1"`
}

func (o *Rfc1123datemap) GetDateTime() map[string]string {
	if o == nil {
		return nil
	}
	return o.DateTime
}

func (o *Rfc1123datemap) GetDateTime1() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.DateTime1
}
