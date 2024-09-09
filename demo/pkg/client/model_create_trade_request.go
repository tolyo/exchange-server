/*
OPEN OUTCRY API

# Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateTradeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeRequest{}

// CreateTradeRequest struct for CreateTradeRequest
type CreateTradeRequest struct {
	// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
	Instrument  string                `json:"instrument"`
	Side        TradeOrderSide        `json:"side"`
	Type        TradeOrderType        `json:"type"`
	TimeInForce TradeOrderTimeInForce `json:"timeInForce"`
	Amount      float64               `json:"amount"`
	Price       *float64              `json:"price,omitempty"`
}

type _CreateTradeRequest CreateTradeRequest

// NewCreateTradeRequest instantiates a new CreateTradeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeRequest(instrument string, side TradeOrderSide, type_ TradeOrderType, timeInForce TradeOrderTimeInForce, amount float64) *CreateTradeRequest {
	this := CreateTradeRequest{}
	this.Instrument = instrument
	this.Side = side
	this.Type = type_
	this.TimeInForce = timeInForce
	this.Amount = amount
	return &this
}

// NewCreateTradeRequestWithDefaults instantiates a new CreateTradeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeRequestWithDefaults() *CreateTradeRequest {
	this := CreateTradeRequest{}
	return &this
}

// GetInstrument returns the Instrument field value
func (o *CreateTradeRequest) GetInstrument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instrument, true
}

// SetInstrument sets field value
func (o *CreateTradeRequest) SetInstrument(v string) {
	o.Instrument = v
}

// GetSide returns the Side field value
func (o *CreateTradeRequest) GetSide() TradeOrderSide {
	if o == nil {
		var ret TradeOrderSide
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetSideOk() (*TradeOrderSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *CreateTradeRequest) SetSide(v TradeOrderSide) {
	o.Side = v
}

// GetType returns the Type field value
func (o *CreateTradeRequest) GetType() TradeOrderType {
	if o == nil {
		var ret TradeOrderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetTypeOk() (*TradeOrderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateTradeRequest) SetType(v TradeOrderType) {
	o.Type = v
}

// GetTimeInForce returns the TimeInForce field value
func (o *CreateTradeRequest) GetTimeInForce() TradeOrderTimeInForce {
	if o == nil {
		var ret TradeOrderTimeInForce
		return ret
	}

	return o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetTimeInForceOk() (*TradeOrderTimeInForce, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInForce, true
}

// SetTimeInForce sets field value
func (o *CreateTradeRequest) SetTimeInForce(v TradeOrderTimeInForce) {
	o.TimeInForce = v
}

// GetAmount returns the Amount field value
func (o *CreateTradeRequest) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateTradeRequest) SetAmount(v float64) {
	o.Amount = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CreateTradeRequest) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeRequest) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CreateTradeRequest) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *CreateTradeRequest) SetPrice(v float64) {
	o.Price = &v
}

func (o CreateTradeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instrument"] = o.Instrument
	toSerialize["side"] = o.Side
	toSerialize["type"] = o.Type
	toSerialize["timeInForce"] = o.TimeInForce
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

func (o *CreateTradeRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instrument",
		"side",
		"type",
		"timeInForce",
		"amount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateTradeRequest := _CreateTradeRequest{}

	err = json.Unmarshal(bytes, &varCreateTradeRequest)

	if err != nil {
		return err
	}

	*o = CreateTradeRequest(varCreateTradeRequest)

	return err
}

type NullableCreateTradeRequest struct {
	value *CreateTradeRequest
	isSet bool
}

func (v NullableCreateTradeRequest) Get() *CreateTradeRequest {
	return v.value
}

func (v *NullableCreateTradeRequest) Set(val *CreateTradeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeRequest(val *CreateTradeRequest) *NullableCreateTradeRequest {
	return &NullableCreateTradeRequest{value: val, isSet: true}
}

func (v NullableCreateTradeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
