/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"fmt"
)

type TradeOrderType string

// List of TradeOrderType
const (
	LIMIT     TradeOrderType = "LIMIT"
	MARKET    TradeOrderType = "MARKET"
	STOPLOSS  TradeOrderType = "STOPLOSS"
	STOPLIMIT TradeOrderType = "STOPLIMIT"
)

// AllowedTradeOrderTypeEnumValues is all the allowed values of TradeOrderType enum
var AllowedTradeOrderTypeEnumValues = []TradeOrderType{
	"LIMIT",
	"MARKET",
	"STOPLOSS",
	"STOPLIMIT",
}

// validTradeOrderTypeEnumValue provides a map of TradeOrderTypes for fast verification of use input
var validTradeOrderTypeEnumValues = map[TradeOrderType]struct{}{
	"LIMIT":     {},
	"MARKET":    {},
	"STOPLOSS":  {},
	"STOPLIMIT": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradeOrderType) IsValid() bool {
	_, ok := validTradeOrderTypeEnumValues[v]
	return ok
}

// NewTradeOrderTypeFromValue returns a pointer to a valid TradeOrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradeOrderTypeFromValue(v string) (TradeOrderType, error) {
	ev := TradeOrderType(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for TradeOrderType: valid values are %v", v, AllowedTradeOrderTypeEnumValues)
	}
}

// AssertTradeOrderTypeRequired checks if the required fields are not zero-ed
func AssertTradeOrderTypeRequired(obj TradeOrderType) error {
	return nil
}

// AssertTradeOrderTypeConstraints checks if the values respects the defined constraints
func AssertTradeOrderTypeConstraints(obj TradeOrderType) error {
	return nil
}
