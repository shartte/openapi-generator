/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// Zebra struct for Zebra
type Zebra struct {
	Type *string `json:"type,omitempty"`
	ClassName string `json:"className"`
}

// NewZebra instantiates a new Zebra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZebra(className string, ) *Zebra {
    this := Zebra{}
    this.ClassName = className
    return &this
}

// NewZebraWithDefaults instantiates a new Zebra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZebraWithDefaults() *Zebra {
    this := Zebra{}
    return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Zebra) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Zebra) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Zebra) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Zebra) SetType(v string) {
	o.Type = &v
}

// GetClassName returns the ClassName field value
func (o *Zebra) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// SetClassName sets field value
func (o *Zebra) SetClassName(v string) {
	o.ClassName = v
}

// AsMammal wraps this instance of Zebra in Mammal
func (s *Zebra) AsMammal() Mammal {
    return Mammal{ MammalInterface: s }
}
type NullableZebra struct {
	Value Zebra
	ExplicitNull bool
}

func (v NullableZebra) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableZebra) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}