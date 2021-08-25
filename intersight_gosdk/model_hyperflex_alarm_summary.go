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
	"reflect"
	"strings"
)

// HyperflexAlarmSummary The summary of alarm counts based on alarm severity.
type HyperflexAlarmSummary struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of alarms that have severity type Critical.
	Critical *int64 `json:"Critical,omitempty"`
	// The count of alarms that have severity type Warning.
	Warning              *int64 `json:"Warning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAlarmSummary HyperflexAlarmSummary

// NewHyperflexAlarmSummary instantiates a new HyperflexAlarmSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAlarmSummary(classId string, objectType string) *HyperflexAlarmSummary {
	this := HyperflexAlarmSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAlarmSummaryWithDefaults instantiates a new HyperflexAlarmSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAlarmSummaryWithDefaults() *HyperflexAlarmSummary {
	this := HyperflexAlarmSummary{}
	var classId string = "hyperflex.AlarmSummary"
	this.ClassId = classId
	var objectType string = "hyperflex.AlarmSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAlarmSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAlarmSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAlarmSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAlarmSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *HyperflexAlarmSummary) GetCritical() int64 {
	if o == nil || o.Critical == nil {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummary) GetCriticalOk() (*int64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *HyperflexAlarmSummary) HasCritical() bool {
	if o != nil && o.Critical != nil {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *HyperflexAlarmSummary) SetCritical(v int64) {
	o.Critical = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *HyperflexAlarmSummary) GetWarning() int64 {
	if o == nil || o.Warning == nil {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAlarmSummary) GetWarningOk() (*int64, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *HyperflexAlarmSummary) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *HyperflexAlarmSummary) SetWarning(v int64) {
	o.Warning = &v
}

func (o HyperflexAlarmSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Critical != nil {
		toSerialize["Critical"] = o.Critical
	}
	if o.Warning != nil {
		toSerialize["Warning"] = o.Warning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAlarmSummary) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexAlarmSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The count of alarms that have severity type Critical.
		Critical *int64 `json:"Critical,omitempty"`
		// The count of alarms that have severity type Warning.
		Warning *int64 `json:"Warning,omitempty"`
	}

	varHyperflexAlarmSummaryWithoutEmbeddedStruct := HyperflexAlarmSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexAlarmSummaryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexAlarmSummary := _HyperflexAlarmSummary{}
		varHyperflexAlarmSummary.ClassId = varHyperflexAlarmSummaryWithoutEmbeddedStruct.ClassId
		varHyperflexAlarmSummary.ObjectType = varHyperflexAlarmSummaryWithoutEmbeddedStruct.ObjectType
		varHyperflexAlarmSummary.Critical = varHyperflexAlarmSummaryWithoutEmbeddedStruct.Critical
		varHyperflexAlarmSummary.Warning = varHyperflexAlarmSummaryWithoutEmbeddedStruct.Warning
		*o = HyperflexAlarmSummary(varHyperflexAlarmSummary)
	} else {
		return err
	}

	varHyperflexAlarmSummary := _HyperflexAlarmSummary{}

	err = json.Unmarshal(bytes, &varHyperflexAlarmSummary)
	if err == nil {
		o.MoBaseComplexType = varHyperflexAlarmSummary.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Critical")
		delete(additionalProperties, "Warning")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAlarmSummary struct {
	value *HyperflexAlarmSummary
	isSet bool
}

func (v NullableHyperflexAlarmSummary) Get() *HyperflexAlarmSummary {
	return v.value
}

func (v *NullableHyperflexAlarmSummary) Set(val *HyperflexAlarmSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAlarmSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAlarmSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAlarmSummary(val *HyperflexAlarmSummary) *NullableHyperflexAlarmSummary {
	return &NullableHyperflexAlarmSummary{value: val, isSet: true}
}

func (v NullableHyperflexAlarmSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAlarmSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
