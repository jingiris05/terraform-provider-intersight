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

// WorkflowPrimitiveDataTypeAllOf Definition of the list of properties defined in 'workflow.PrimitiveDataType', excluding properties defined in parent classes.
type WorkflowPrimitiveDataTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	Properties           NullableWorkflowPrimitiveDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveDataTypeAllOf WorkflowPrimitiveDataTypeAllOf

// NewWorkflowPrimitiveDataTypeAllOf instantiates a new WorkflowPrimitiveDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveDataTypeAllOf(classId string, objectType string) *WorkflowPrimitiveDataTypeAllOf {
	this := WorkflowPrimitiveDataTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPrimitiveDataTypeAllOfWithDefaults instantiates a new WorkflowPrimitiveDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveDataTypeAllOfWithDefaults() *WorkflowPrimitiveDataTypeAllOf {
	this := WorkflowPrimitiveDataTypeAllOf{}
	var classId string = "workflow.PrimitiveDataType"
	this.ClassId = classId
	var objectType string = "workflow.PrimitiveDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPrimitiveDataTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPrimitiveDataTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPrimitiveDataTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPrimitiveDataTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPrimitiveDataTypeAllOf) GetProperties() WorkflowPrimitiveDataProperty {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowPrimitiveDataProperty
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPrimitiveDataTypeAllOf) GetPropertiesOk() (*WorkflowPrimitiveDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataTypeAllOf) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowPrimitiveDataProperty and assigns it to the Properties field.
func (o *WorkflowPrimitiveDataTypeAllOf) SetProperties(v WorkflowPrimitiveDataProperty) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowPrimitiveDataTypeAllOf) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowPrimitiveDataTypeAllOf) UnsetProperties() {
	o.Properties.Unset()
}

func (o WorkflowPrimitiveDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Properties.IsSet() {
		toSerialize["Properties"] = o.Properties.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowPrimitiveDataTypeAllOf := _WorkflowPrimitiveDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataTypeAllOf); err == nil {
		*o = WorkflowPrimitiveDataTypeAllOf(varWorkflowPrimitiveDataTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPrimitiveDataTypeAllOf struct {
	value *WorkflowPrimitiveDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowPrimitiveDataTypeAllOf) Get() *WorkflowPrimitiveDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowPrimitiveDataTypeAllOf) Set(val *WorkflowPrimitiveDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPrimitiveDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPrimitiveDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPrimitiveDataTypeAllOf(val *WorkflowPrimitiveDataTypeAllOf) *NullableWorkflowPrimitiveDataTypeAllOf {
	return &NullableWorkflowPrimitiveDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowPrimitiveDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPrimitiveDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
