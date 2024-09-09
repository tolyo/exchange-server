/*
OPEN OUTCRY API

# Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PriceVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceVolume{}

// PriceVolume struct for PriceVolume
type PriceVolume struct {
	Price  *float32 `json:"price,omitempty"`
	Volume *float32 `json:"volume,omitempty"`
}

// NewPriceVolume instantiates a new PriceVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolume() *PriceVolume {
	this := PriceVolume{}
	return &this
}

// NewPriceVolumeWithDefaults instantiates a new PriceVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeWithDefaults() *PriceVolume {
	this := PriceVolume{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceVolume) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolume) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceVolume) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PriceVolume) SetPrice(v float32) {
	o.Price = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *PriceVolume) GetVolume() float32 {
	if o == nil || IsNil(o.Volume) {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolume) GetVolumeOk() (*float32, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *PriceVolume) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *PriceVolume) SetVolume(v float32) {
	o.Volume = &v
}

func (o PriceVolume) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullablePriceVolume struct {
	value *PriceVolume
	isSet bool
}

func (v NullablePriceVolume) Get() *PriceVolume {
	return v.value
}

func (v *NullablePriceVolume) Set(val *PriceVolume) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolume) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolume(val *PriceVolume) *NullablePriceVolume {
	return &NullablePriceVolume{value: val, isSet: true}
}

func (v NullablePriceVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
