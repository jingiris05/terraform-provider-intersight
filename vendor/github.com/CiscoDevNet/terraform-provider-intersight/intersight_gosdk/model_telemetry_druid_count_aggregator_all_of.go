/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidCountAggregatorAllOf struct for TelemetryDruidCountAggregatorAllOf
type TelemetryDruidCountAggregatorAllOf struct {
	// the output name
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidCountAggregatorAllOf TelemetryDruidCountAggregatorAllOf

// NewTelemetryDruidCountAggregatorAllOf instantiates a new TelemetryDruidCountAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidCountAggregatorAllOf() *TelemetryDruidCountAggregatorAllOf {
	this := TelemetryDruidCountAggregatorAllOf{}
	return &this
}

// NewTelemetryDruidCountAggregatorAllOfWithDefaults instantiates a new TelemetryDruidCountAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidCountAggregatorAllOfWithDefaults() *TelemetryDruidCountAggregatorAllOf {
	this := TelemetryDruidCountAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidCountAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidCountAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidCountAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidCountAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

func (o TelemetryDruidCountAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidCountAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidCountAggregatorAllOf := _TelemetryDruidCountAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidCountAggregatorAllOf); err == nil {
		*o = TelemetryDruidCountAggregatorAllOf(varTelemetryDruidCountAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidCountAggregatorAllOf struct {
	value *TelemetryDruidCountAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidCountAggregatorAllOf) Get() *TelemetryDruidCountAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidCountAggregatorAllOf) Set(val *TelemetryDruidCountAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidCountAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidCountAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidCountAggregatorAllOf(val *TelemetryDruidCountAggregatorAllOf) *NullableTelemetryDruidCountAggregatorAllOf {
	return &NullableTelemetryDruidCountAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidCountAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidCountAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
