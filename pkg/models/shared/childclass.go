// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChildClass struct {
	BigDecimal                             *string      `json:"Big_Decimal,omitempty"`
	ChildClassArray                        []ChildClass `json:"Child_Class_Array,omitempty"`
	GrandParentOptional                    *string      `json:"Grand_Parent_Optional,omitempty"`
	GrandParentRequired                    string       `json:"Grand_Parent_Required"`
	GrandParentRequiredNullable            string       `json:"Grand_Parent_Required_Nullable"`
	Optional                               *string      `json:"Optional,omitempty"`
	OptionalNullable                       *string      `json:"Optional_Nullable,omitempty"`
	OptionalNullableWithDefaultValue       *string      `json:"Optional_Nullable_With_Default_Value,omitempty"`
	ParentOptional                         *string      `json:"Parent_Optional,omitempty"`
	ParentOptionalNullableWithDefaultValue *string      `json:"Parent_Optional_Nullable_With_Default_Value,omitempty"`
	ParentRequired                         string       `json:"Parent_Required"`
	ParentRequiredNullable                 string       `json:"Parent_Required_Nullable"`
	Required                               string       `json:"Required"`
	RequiredNullable                       string       `json:"Required_Nullable"`
	Class                                  *int         `json:"class,omitempty"`
	Precision                              *float64     `json:"precision,omitempty"`
}

func (o *ChildClass) GetBigDecimal() *string {
	if o == nil {
		return nil
	}
	return o.BigDecimal
}

func (o *ChildClass) GetChildClassArray() []ChildClass {
	if o == nil {
		return nil
	}
	return o.ChildClassArray
}

func (o *ChildClass) GetGrandParentOptional() *string {
	if o == nil {
		return nil
	}
	return o.GrandParentOptional
}

func (o *ChildClass) GetGrandParentRequired() string {
	if o == nil {
		return ""
	}
	return o.GrandParentRequired
}

func (o *ChildClass) GetGrandParentRequiredNullable() string {
	if o == nil {
		return ""
	}
	return o.GrandParentRequiredNullable
}

func (o *ChildClass) GetOptional() *string {
	if o == nil {
		return nil
	}
	return o.Optional
}

func (o *ChildClass) GetOptionalNullable() *string {
	if o == nil {
		return nil
	}
	return o.OptionalNullable
}

func (o *ChildClass) GetOptionalNullableWithDefaultValue() *string {
	if o == nil {
		return nil
	}
	return o.OptionalNullableWithDefaultValue
}

func (o *ChildClass) GetParentOptional() *string {
	if o == nil {
		return nil
	}
	return o.ParentOptional
}

func (o *ChildClass) GetParentOptionalNullableWithDefaultValue() *string {
	if o == nil {
		return nil
	}
	return o.ParentOptionalNullableWithDefaultValue
}

func (o *ChildClass) GetParentRequired() string {
	if o == nil {
		return ""
	}
	return o.ParentRequired
}

func (o *ChildClass) GetParentRequiredNullable() string {
	if o == nil {
		return ""
	}
	return o.ParentRequiredNullable
}

func (o *ChildClass) GetRequired() string {
	if o == nil {
		return ""
	}
	return o.Required
}

func (o *ChildClass) GetRequiredNullable() string {
	if o == nil {
		return ""
	}
	return o.RequiredNullable
}

func (o *ChildClass) GetClass() *int {
	if o == nil {
		return nil
	}
	return o.Class
}

func (o *ChildClass) GetPrecision() *float64 {
	if o == nil {
		return nil
	}
	return o.Precision
}