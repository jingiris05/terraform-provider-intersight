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

// checks if the TelemetryDruidInFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidInFilter{}

// TelemetryDruidInFilter The in filter can match input rows against a set of values, where a match occurs if the value is contained in the set. If an empty values array is passed to the \"in\" filter, it will simply return an empty result. If the values array contains null, the \"in\" filter matches null values.
type TelemetryDruidInFilter struct {
	Type string `json:"type"`
	// Input column or virtual column name to filter.
	Dimension string `json:"dimension"`
	// List of string value to match.
	Values               []string                          `json:"values"`
	ExtractionFn         *TelemetryDruidExtractionFunction `json:"extractionFn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInFilter TelemetryDruidInFilter

// NewTelemetryDruidInFilter instantiates a new TelemetryDruidInFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInFilter(type_ string, dimension string, values []string) *TelemetryDruidInFilter {
	this := TelemetryDruidInFilter{}
	this.Type = type_
	this.Dimension = dimension
	this.Values = values
	return &this
}

// NewTelemetryDruidInFilterWithDefaults instantiates a new TelemetryDruidInFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInFilterWithDefaults() *TelemetryDruidInFilter {
	this := TelemetryDruidInFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidInFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidInFilter) SetType(v string) {
	o.Type = v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidInFilter) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInFilter) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidInFilter) SetDimension(v string) {
	o.Dimension = v
}

// GetValues returns the Values field value
func (o *TelemetryDruidInFilter) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInFilter) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *TelemetryDruidInFilter) SetValues(v []string) {
	o.Values = v
}

// GetExtractionFn returns the ExtractionFn field value if set, zero value otherwise.
func (o *TelemetryDruidInFilter) GetExtractionFn() TelemetryDruidExtractionFunction {
	if o == nil || IsNil(o.ExtractionFn) {
		var ret TelemetryDruidExtractionFunction
		return ret
	}
	return *o.ExtractionFn
}

// GetExtractionFnOk returns a tuple with the ExtractionFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInFilter) GetExtractionFnOk() (*TelemetryDruidExtractionFunction, bool) {
	if o == nil || IsNil(o.ExtractionFn) {
		return nil, false
	}
	return o.ExtractionFn, true
}

// HasExtractionFn returns a boolean if a field has been set.
func (o *TelemetryDruidInFilter) HasExtractionFn() bool {
	if o != nil && !IsNil(o.ExtractionFn) {
		return true
	}

	return false
}

// SetExtractionFn gets a reference to the given TelemetryDruidExtractionFunction and assigns it to the ExtractionFn field.
func (o *TelemetryDruidInFilter) SetExtractionFn(v TelemetryDruidExtractionFunction) {
	o.ExtractionFn = &v
}

func (o TelemetryDruidInFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidInFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["dimension"] = o.Dimension
	toSerialize["values"] = o.Values
	if !IsNil(o.ExtractionFn) {
		toSerialize["extractionFn"] = o.ExtractionFn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidInFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"dimension",
		"values",
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
	varTelemetryDruidInFilter := _TelemetryDruidInFilter{}

	err = json.Unmarshal(data, &varTelemetryDruidInFilter)

	if err != nil {
		return err
	}

	*o = TelemetryDruidInFilter(varTelemetryDruidInFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "dimension")
		delete(additionalProperties, "values")
		delete(additionalProperties, "extractionFn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInFilter struct {
	value *TelemetryDruidInFilter
	isSet bool
}

func (v NullableTelemetryDruidInFilter) Get() *TelemetryDruidInFilter {
	return v.value
}

func (v *NullableTelemetryDruidInFilter) Set(val *TelemetryDruidInFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInFilter(val *TelemetryDruidInFilter) *NullableTelemetryDruidInFilter {
	return &NullableTelemetryDruidInFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidInFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
