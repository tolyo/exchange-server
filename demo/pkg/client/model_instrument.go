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

// checks if the Instrument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instrument{}

// Instrument struct for Instrument
type Instrument struct {
	Id *string `json:"id,omitempty"`
	// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
	Name *string `json:"name,omitempty"`
	// ISO 4217 Currency symbol
	QuoteCurrency *string `json:"quote_currency,omitempty"`
	// Availability for trading
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInstrument instantiates a new Instrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstrument() *Instrument {
	this := Instrument{}
	return &this
}

// NewInstrumentWithDefaults instantiates a new Instrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstrumentWithDefaults() *Instrument {
	this := Instrument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Instrument) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Instrument) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Instrument) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Instrument) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Instrument) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Instrument) SetName(v string) {
	o.Name = &v
}

// GetQuoteCurrency returns the QuoteCurrency field value if set, zero value otherwise.
func (o *Instrument) GetQuoteCurrency() string {
	if o == nil || IsNil(o.QuoteCurrency) {
		var ret string
		return ret
	}
	return *o.QuoteCurrency
}

// GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetQuoteCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteCurrency) {
		return nil, false
	}
	return o.QuoteCurrency, true
}

// HasQuoteCurrency returns a boolean if a field has been set.
func (o *Instrument) HasQuoteCurrency() bool {
	if o != nil && !IsNil(o.QuoteCurrency) {
		return true
	}

	return false
}

// SetQuoteCurrency gets a reference to the given string and assigns it to the QuoteCurrency field.
func (o *Instrument) SetQuoteCurrency(v string) {
	o.QuoteCurrency = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Instrument) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Instrument) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Instrument) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o Instrument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.QuoteCurrency) {
		toSerialize["quote_currency"] = o.QuoteCurrency
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableInstrument struct {
	value *Instrument
	isSet bool
}

func (v NullableInstrument) Get() *Instrument {
	return v.value
}

func (v *NullableInstrument) Set(val *Instrument) {
	v.value = val
	v.isSet = true
}

func (v NullableInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstrument(val *Instrument) *NullableInstrument {
	return &NullableInstrument{value: val, isSet: true}
}

func (v NullableInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
