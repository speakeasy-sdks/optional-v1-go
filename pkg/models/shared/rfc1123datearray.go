// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Rfc1123datearray struct {
	DateTime  []string `json:"dateTime,omitempty"`
	DateTime1 []string `json:"dateTime1"`
}

func (o *Rfc1123datearray) GetDateTime() []string {
	if o == nil {
		return nil
	}
	return o.DateTime
}

func (o *Rfc1123datearray) GetDateTime1() []string {
	if o == nil {
		return []string{}
	}
	return o.DateTime1
}
