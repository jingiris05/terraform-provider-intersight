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

// checks if the TelemetryDruidFirstLastAggregator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidFirstLastAggregator{}

// TelemetryDruidFirstLastAggregator (Double/Float/Long) First and Last aggregator cannot be used in ingestion spec, and should only be specified as part of queries. Note that queries with first/last aggregators on a segment created with rollup enabled will return the rolled up value, and not the last value within the raw ingested data.
type TelemetryDruidFirstLastAggregator struct {
	// The aggregator type.
	Type string `json:"type"`
	// Output name for the first/last value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidFirstLastAggregator TelemetryDruidFirstLastAggregator

// NewTelemetryDruidFirstLastAggregator instantiates a new TelemetryDruidFirstLastAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidFirstLastAggregator(type_ string, name string, fieldName string) *TelemetryDruidFirstLastAggregator {
	this := TelemetryDruidFirstLastAggregator{}
	this.Type = type_
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidFirstLastAggregatorWithDefaults instantiates a new TelemetryDruidFirstLastAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidFirstLastAggregatorWithDefaults() *TelemetryDruidFirstLastAggregator {
	this := TelemetryDruidFirstLastAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidFirstLastAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFirstLastAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidFirstLastAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TelemetryDruidFirstLastAggregator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFirstLastAggregator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidFirstLastAggregator) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidFirstLastAggregator) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFirstLastAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidFirstLastAggregator) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidFirstLastAggregator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidFirstLastAggregator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["fieldName"] = o.FieldName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidFirstLastAggregator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"fieldName",
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
	varTelemetryDruidFirstLastAggregator := _TelemetryDruidFirstLastAggregator{}

	err = json.Unmarshal(data, &varTelemetryDruidFirstLastAggregator)

	if err != nil {
		return err
	}

	*o = TelemetryDruidFirstLastAggregator(varTelemetryDruidFirstLastAggregator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidFirstLastAggregator struct {
	value *TelemetryDruidFirstLastAggregator
	isSet bool
}

func (v NullableTelemetryDruidFirstLastAggregator) Get() *TelemetryDruidFirstLastAggregator {
	return v.value
}

func (v *NullableTelemetryDruidFirstLastAggregator) Set(val *TelemetryDruidFirstLastAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFirstLastAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFirstLastAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFirstLastAggregator(val *TelemetryDruidFirstLastAggregator) *NullableTelemetryDruidFirstLastAggregator {
	return &NullableTelemetryDruidFirstLastAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidFirstLastAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFirstLastAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
