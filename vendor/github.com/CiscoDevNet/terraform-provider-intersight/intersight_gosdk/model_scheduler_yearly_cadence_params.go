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
	"reflect"
	"strings"
)

// checks if the SchedulerYearlyCadenceParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerYearlyCadenceParams{}

// SchedulerYearlyCadenceParams Parameters for a yearly cadence.
type SchedulerYearlyCadenceParams struct {
	SchedulerBaseMonthlyCadenceParams
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	MonthOfYear          []string `json:"MonthOfYear,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerYearlyCadenceParams SchedulerYearlyCadenceParams

// NewSchedulerYearlyCadenceParams instantiates a new SchedulerYearlyCadenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerYearlyCadenceParams(classId string, objectType string) *SchedulerYearlyCadenceParams {
	this := SchedulerYearlyCadenceParams{}
	this.ClassId = classId
	this.ObjectType = objectType
	var customDayOfMonth string = "None"
	this.CustomDayOfMonth = &customDayOfMonth
	return &this
}

// NewSchedulerYearlyCadenceParamsWithDefaults instantiates a new SchedulerYearlyCadenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerYearlyCadenceParamsWithDefaults() *SchedulerYearlyCadenceParams {
	this := SchedulerYearlyCadenceParams{}
	var classId string = "scheduler.YearlyCadenceParams"
	this.ClassId = classId
	var objectType string = "scheduler.YearlyCadenceParams"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerYearlyCadenceParams) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerYearlyCadenceParams) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerYearlyCadenceParams) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.YearlyCadenceParams" of the ClassId field.
func (o *SchedulerYearlyCadenceParams) GetDefaultClassId() interface{} {
	return "scheduler.YearlyCadenceParams"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerYearlyCadenceParams) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerYearlyCadenceParams) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerYearlyCadenceParams) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.YearlyCadenceParams" of the ObjectType field.
func (o *SchedulerYearlyCadenceParams) GetDefaultObjectType() interface{} {
	return "scheduler.YearlyCadenceParams"
}

// GetMonthOfYear returns the MonthOfYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulerYearlyCadenceParams) GetMonthOfYear() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MonthOfYear
}

// GetMonthOfYearOk returns a tuple with the MonthOfYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulerYearlyCadenceParams) GetMonthOfYearOk() ([]string, bool) {
	if o == nil || IsNil(o.MonthOfYear) {
		return nil, false
	}
	return o.MonthOfYear, true
}

// HasMonthOfYear returns a boolean if a field has been set.
func (o *SchedulerYearlyCadenceParams) HasMonthOfYear() bool {
	if o != nil && !IsNil(o.MonthOfYear) {
		return true
	}

	return false
}

// SetMonthOfYear gets a reference to the given []string and assigns it to the MonthOfYear field.
func (o *SchedulerYearlyCadenceParams) SetMonthOfYear(v []string) {
	o.MonthOfYear = v
}

func (o SchedulerYearlyCadenceParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerYearlyCadenceParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSchedulerBaseMonthlyCadenceParams, errSchedulerBaseMonthlyCadenceParams := json.Marshal(o.SchedulerBaseMonthlyCadenceParams)
	if errSchedulerBaseMonthlyCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseMonthlyCadenceParams
	}
	errSchedulerBaseMonthlyCadenceParams = json.Unmarshal([]byte(serializedSchedulerBaseMonthlyCadenceParams), &toSerialize)
	if errSchedulerBaseMonthlyCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseMonthlyCadenceParams
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.MonthOfYear != nil {
		toSerialize["MonthOfYear"] = o.MonthOfYear
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerYearlyCadenceParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
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
	type SchedulerYearlyCadenceParamsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string   `json:"ObjectType"`
		MonthOfYear []string `json:"MonthOfYear,omitempty"`
	}

	varSchedulerYearlyCadenceParamsWithoutEmbeddedStruct := SchedulerYearlyCadenceParamsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerYearlyCadenceParamsWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerYearlyCadenceParams := _SchedulerYearlyCadenceParams{}
		varSchedulerYearlyCadenceParams.ClassId = varSchedulerYearlyCadenceParamsWithoutEmbeddedStruct.ClassId
		varSchedulerYearlyCadenceParams.ObjectType = varSchedulerYearlyCadenceParamsWithoutEmbeddedStruct.ObjectType
		varSchedulerYearlyCadenceParams.MonthOfYear = varSchedulerYearlyCadenceParamsWithoutEmbeddedStruct.MonthOfYear
		*o = SchedulerYearlyCadenceParams(varSchedulerYearlyCadenceParams)
	} else {
		return err
	}

	varSchedulerYearlyCadenceParams := _SchedulerYearlyCadenceParams{}

	err = json.Unmarshal(data, &varSchedulerYearlyCadenceParams)
	if err == nil {
		o.SchedulerBaseMonthlyCadenceParams = varSchedulerYearlyCadenceParams.SchedulerBaseMonthlyCadenceParams
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MonthOfYear")

		// remove fields from embedded structs
		reflectSchedulerBaseMonthlyCadenceParams := reflect.ValueOf(o.SchedulerBaseMonthlyCadenceParams)
		for i := 0; i < reflectSchedulerBaseMonthlyCadenceParams.Type().NumField(); i++ {
			t := reflectSchedulerBaseMonthlyCadenceParams.Type().Field(i)

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

type NullableSchedulerYearlyCadenceParams struct {
	value *SchedulerYearlyCadenceParams
	isSet bool
}

func (v NullableSchedulerYearlyCadenceParams) Get() *SchedulerYearlyCadenceParams {
	return v.value
}

func (v *NullableSchedulerYearlyCadenceParams) Set(val *SchedulerYearlyCadenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerYearlyCadenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerYearlyCadenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerYearlyCadenceParams(val *SchedulerYearlyCadenceParams) *NullableSchedulerYearlyCadenceParams {
	return &NullableSchedulerYearlyCadenceParams{value: val, isSet: true}
}

func (v NullableSchedulerYearlyCadenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerYearlyCadenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
