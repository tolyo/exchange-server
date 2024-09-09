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

// checks if the TradingAccountInstrument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradingAccountInstrument{}

// TradingAccountInstrument struct for TradingAccountInstrument
type TradingAccountInstrument struct {
	// Ticker-like name of the instrument. For monetary instruments, a currency pair is used.
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	AmountReserved  float64 `json:"amountReserved"`
	AmountAvailable float64 `json:"amountAvailable"`
	Value           float64 `json:"value"`
	// ISO 4217 Currency symbol
	Currency string `json:"currency"`
}

type _TradingAccountInstrument TradingAccountInstrument

// NewTradingAccountInstrument instantiates a new TradingAccountInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingAccountInstrument(name string, amount float64, amountReserved float64, amountAvailable float64, value float64, currency string) *TradingAccountInstrument {
	this := TradingAccountInstrument{}
	this.Name = name
	this.Amount = amount
	this.AmountReserved = amountReserved
	this.AmountAvailable = amountAvailable
	this.Value = value
	this.Currency = currency
	return &this
}

// NewTradingAccountInstrumentWithDefaults instantiates a new TradingAccountInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingAccountInstrumentWithDefaults() *TradingAccountInstrument {
	this := TradingAccountInstrument{}
	return &this
}

// GetName returns the Name field value
func (o *TradingAccountInstrument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TradingAccountInstrument) SetName(v string) {
	o.Name = v
}

// GetAmount returns the Amount field value
func (o *TradingAccountInstrument) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TradingAccountInstrument) SetAmount(v float64) {
	o.Amount = v
}

// GetAmountReserved returns the AmountReserved field value
func (o *TradingAccountInstrument) GetAmountReserved() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountReserved
}

// GetAmountReservedOk returns a tuple with the AmountReserved field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetAmountReservedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountReserved, true
}

// SetAmountReserved sets field value
func (o *TradingAccountInstrument) SetAmountReserved(v float64) {
	o.AmountReserved = v
}

// GetAmountAvailable returns the AmountAvailable field value
func (o *TradingAccountInstrument) GetAmountAvailable() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountAvailable
}

// GetAmountAvailableOk returns a tuple with the AmountAvailable field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetAmountAvailableOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAvailable, true
}

// SetAmountAvailable sets field value
func (o *TradingAccountInstrument) SetAmountAvailable(v float64) {
	o.AmountAvailable = v
}

// GetValue returns the Value field value
func (o *TradingAccountInstrument) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TradingAccountInstrument) SetValue(v float64) {
	o.Value = v
}

// GetCurrency returns the Currency field value
func (o *TradingAccountInstrument) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TradingAccountInstrument) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TradingAccountInstrument) SetCurrency(v string) {
	o.Currency = v
}

func (o TradingAccountInstrument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingAccountInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["amount"] = o.Amount
	toSerialize["amountReserved"] = o.AmountReserved
	toSerialize["amountAvailable"] = o.AmountAvailable
	toSerialize["value"] = o.Value
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *TradingAccountInstrument) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"amount",
		"amountReserved",
		"amountAvailable",
		"value",
		"currency",
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

	varTradingAccountInstrument := _TradingAccountInstrument{}

	err = json.Unmarshal(bytes, &varTradingAccountInstrument)

	if err != nil {
		return err
	}

	*o = TradingAccountInstrument(varTradingAccountInstrument)

	return err
}

type NullableTradingAccountInstrument struct {
	value *TradingAccountInstrument
	isSet bool
}

func (v NullableTradingAccountInstrument) Get() *TradingAccountInstrument {
	return v.value
}

func (v *NullableTradingAccountInstrument) Set(val *TradingAccountInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingAccountInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingAccountInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingAccountInstrument(val *TradingAccountInstrument) *NullableTradingAccountInstrument {
	return &NullableTradingAccountInstrument{value: val, isSet: true}
}

func (v NullableTradingAccountInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingAccountInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
