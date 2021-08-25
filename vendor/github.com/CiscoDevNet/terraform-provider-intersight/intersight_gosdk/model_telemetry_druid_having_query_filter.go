/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidHavingQueryFilter Query filter HavingSpecs allow all Druid query filters to be used in the Having part of the query.
type TelemetryDruidHavingQueryFilter struct {
	// The having filter type.
	Type                 string               `json:"type"`
	Filter               TelemetryDruidFilter `json:"filter"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidHavingQueryFilter TelemetryDruidHavingQueryFilter

// NewTelemetryDruidHavingQueryFilter instantiates a new TelemetryDruidHavingQueryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidHavingQueryFilter(type_ string, filter TelemetryDruidFilter) *TelemetryDruidHavingQueryFilter {
	this := TelemetryDruidHavingQueryFilter{}
	this.Type = type_
	this.Filter = filter
	return &this
}

// NewTelemetryDruidHavingQueryFilterWithDefaults instantiates a new TelemetryDruidHavingQueryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidHavingQueryFilterWithDefaults() *TelemetryDruidHavingQueryFilter {
	this := TelemetryDruidHavingQueryFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidHavingQueryFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingQueryFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidHavingQueryFilter) SetType(v string) {
	o.Type = v
}

// GetFilter returns the Filter field value
func (o *TelemetryDruidHavingQueryFilter) GetFilter() TelemetryDruidFilter {
	if o == nil {
		var ret TelemetryDruidFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingQueryFilter) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *TelemetryDruidHavingQueryFilter) SetFilter(v TelemetryDruidFilter) {
	o.Filter = v
}

func (o TelemetryDruidHavingQueryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidHavingQueryFilter) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidHavingQueryFilter := _TelemetryDruidHavingQueryFilter{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidHavingQueryFilter); err == nil {
		*o = TelemetryDruidHavingQueryFilter(varTelemetryDruidHavingQueryFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidHavingQueryFilter struct {
	value *TelemetryDruidHavingQueryFilter
	isSet bool
}

func (v NullableTelemetryDruidHavingQueryFilter) Get() *TelemetryDruidHavingQueryFilter {
	return v.value
}

func (v *NullableTelemetryDruidHavingQueryFilter) Set(val *TelemetryDruidHavingQueryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHavingQueryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHavingQueryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHavingQueryFilter(val *TelemetryDruidHavingQueryFilter) *NullableTelemetryDruidHavingQueryFilter {
	return &NullableTelemetryDruidHavingQueryFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidHavingQueryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHavingQueryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
