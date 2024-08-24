/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// checks if the TelemetryDruidExtractionFunctionSubstring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidExtractionFunctionSubstring{}

// TelemetryDruidExtractionFunctionSubstring Returns a substring of the dimension value starting from the supplied index and of the desired length. Both index and length are measured in the number of Unicode code units present in the string as if it were encoded in UTF-16. Note that some Unicode characters may be represented by two code units. If the desired length exceeds the length of the dimension value, the remainder of the string starting at index will be returned. If index is greater than the length of the dimension value, null will be returned.
type TelemetryDruidExtractionFunctionSubstring struct {
	Type  string `json:"type"`
	Index int32  `json:"index"`
	// The length may be omitted for substring to return the remainder of the dimension value starting from index, or null if index greater than the length of the dimension value.
	Length               *int32 `json:"length,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionSubstring TelemetryDruidExtractionFunctionSubstring

// NewTelemetryDruidExtractionFunctionSubstring instantiates a new TelemetryDruidExtractionFunctionSubstring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionSubstring(type_ string, index int32) *TelemetryDruidExtractionFunctionSubstring {
	this := TelemetryDruidExtractionFunctionSubstring{}
	this.Type = type_
	this.Index = index
	return &this
}

// NewTelemetryDruidExtractionFunctionSubstringWithDefaults instantiates a new TelemetryDruidExtractionFunctionSubstring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionSubstringWithDefaults() *TelemetryDruidExtractionFunctionSubstring {
	this := TelemetryDruidExtractionFunctionSubstring{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionSubstring) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionSubstring) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionSubstring) SetType(v string) {
	o.Type = v
}

// GetIndex returns the Index field value
func (o *TelemetryDruidExtractionFunctionSubstring) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionSubstring) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *TelemetryDruidExtractionFunctionSubstring) SetIndex(v int32) {
	o.Index = v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionSubstring) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionSubstring) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionSubstring) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *TelemetryDruidExtractionFunctionSubstring) SetLength(v int32) {
	o.Length = &v
}

func (o TelemetryDruidExtractionFunctionSubstring) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidExtractionFunctionSubstring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["index"] = o.Index
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidExtractionFunctionSubstring) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"index",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varTelemetryDruidExtractionFunctionSubstring := _TelemetryDruidExtractionFunctionSubstring{}

	err = json.Unmarshal(data, &varTelemetryDruidExtractionFunctionSubstring)

	if err != nil {
		return err
	}

	*o = TelemetryDruidExtractionFunctionSubstring(varTelemetryDruidExtractionFunctionSubstring)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "index")
		delete(additionalProperties, "length")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionSubstring struct {
	value *TelemetryDruidExtractionFunctionSubstring
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionSubstring) Get() *TelemetryDruidExtractionFunctionSubstring {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionSubstring) Set(val *TelemetryDruidExtractionFunctionSubstring) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionSubstring) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionSubstring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionSubstring(val *TelemetryDruidExtractionFunctionSubstring) *NullableTelemetryDruidExtractionFunctionSubstring {
	return &NullableTelemetryDruidExtractionFunctionSubstring{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionSubstring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionSubstring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
