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

// WorkflowConstraintsAllOf Definition of the list of properties defined in 'workflow.Constraints', excluding properties defined in parent classes.
type WorkflowConstraintsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string              `json:"ObjectType"`
	EnumList   []WorkflowEnumEntry `json:"EnumList,omitempty"`
	// Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.
	Max *float64 `json:"Max,omitempty"`
	// Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.
	Min *float64 `json:"Min,omitempty"`
	// When the parameter is a string this regular expression is used to ensure the value is valid.
	Regex                *string `json:"Regex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowConstraintsAllOf WorkflowConstraintsAllOf

// NewWorkflowConstraintsAllOf instantiates a new WorkflowConstraintsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowConstraintsAllOf(classId string, objectType string) *WorkflowConstraintsAllOf {
	this := WorkflowConstraintsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowConstraintsAllOfWithDefaults instantiates a new WorkflowConstraintsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowConstraintsAllOfWithDefaults() *WorkflowConstraintsAllOf {
	this := WorkflowConstraintsAllOf{}
	var classId string = "workflow.Constraints"
	this.ClassId = classId
	var objectType string = "workflow.Constraints"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowConstraintsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowConstraintsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowConstraintsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowConstraintsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowConstraintsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowConstraintsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnumList returns the EnumList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowConstraintsAllOf) GetEnumList() []WorkflowEnumEntry {
	if o == nil {
		var ret []WorkflowEnumEntry
		return ret
	}
	return o.EnumList
}

// GetEnumListOk returns a tuple with the EnumList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowConstraintsAllOf) GetEnumListOk() (*[]WorkflowEnumEntry, bool) {
	if o == nil || o.EnumList == nil {
		return nil, false
	}
	return &o.EnumList, true
}

// HasEnumList returns a boolean if a field has been set.
func (o *WorkflowConstraintsAllOf) HasEnumList() bool {
	if o != nil && o.EnumList != nil {
		return true
	}

	return false
}

// SetEnumList gets a reference to the given []WorkflowEnumEntry and assigns it to the EnumList field.
func (o *WorkflowConstraintsAllOf) SetEnumList(v []WorkflowEnumEntry) {
	o.EnumList = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *WorkflowConstraintsAllOf) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraintsAllOf) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *WorkflowConstraintsAllOf) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *WorkflowConstraintsAllOf) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *WorkflowConstraintsAllOf) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraintsAllOf) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *WorkflowConstraintsAllOf) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *WorkflowConstraintsAllOf) SetMin(v float64) {
	o.Min = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *WorkflowConstraintsAllOf) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConstraintsAllOf) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *WorkflowConstraintsAllOf) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *WorkflowConstraintsAllOf) SetRegex(v string) {
	o.Regex = &v
}

func (o WorkflowConstraintsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnumList != nil {
		toSerialize["EnumList"] = o.EnumList
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}
	if o.Regex != nil {
		toSerialize["Regex"] = o.Regex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowConstraintsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowConstraintsAllOf := _WorkflowConstraintsAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowConstraintsAllOf); err == nil {
		*o = WorkflowConstraintsAllOf(varWorkflowConstraintsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnumList")
		delete(additionalProperties, "Max")
		delete(additionalProperties, "Min")
		delete(additionalProperties, "Regex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowConstraintsAllOf struct {
	value *WorkflowConstraintsAllOf
	isSet bool
}

func (v NullableWorkflowConstraintsAllOf) Get() *WorkflowConstraintsAllOf {
	return v.value
}

func (v *NullableWorkflowConstraintsAllOf) Set(val *WorkflowConstraintsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowConstraintsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowConstraintsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowConstraintsAllOf(val *WorkflowConstraintsAllOf) *NullableWorkflowConstraintsAllOf {
	return &NullableWorkflowConstraintsAllOf{value: val, isSet: true}
}

func (v NullableWorkflowConstraintsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowConstraintsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
