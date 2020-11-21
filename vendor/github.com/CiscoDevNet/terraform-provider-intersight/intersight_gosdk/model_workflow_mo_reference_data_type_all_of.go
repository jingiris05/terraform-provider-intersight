/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowMoReferenceDataTypeAllOf Definition of the list of properties defined in 'workflow.MoReferenceDataType', excluding properties defined in parent classes.
type WorkflowMoReferenceDataTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                        `json:"ObjectType"`
	Properties           []WorkflowMoReferenceProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMoReferenceDataTypeAllOf WorkflowMoReferenceDataTypeAllOf

// NewWorkflowMoReferenceDataTypeAllOf instantiates a new WorkflowMoReferenceDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoReferenceDataTypeAllOf(classId string, objectType string) *WorkflowMoReferenceDataTypeAllOf {
	this := WorkflowMoReferenceDataTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMoReferenceDataTypeAllOfWithDefaults instantiates a new WorkflowMoReferenceDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoReferenceDataTypeAllOfWithDefaults() *WorkflowMoReferenceDataTypeAllOf {
	this := WorkflowMoReferenceDataTypeAllOf{}
	var classId string = "workflow.MoReferenceDataType"
	this.ClassId = classId
	var objectType string = "workflow.MoReferenceDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoReferenceDataTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceDataTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoReferenceDataTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoReferenceDataTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceDataTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoReferenceDataTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceDataTypeAllOf) GetProperties() []WorkflowMoReferenceProperty {
	if o == nil {
		var ret []WorkflowMoReferenceProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceDataTypeAllOf) GetPropertiesOk() (*[]WorkflowMoReferenceProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowMoReferenceDataTypeAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowMoReferenceProperty and assigns it to the Properties field.
func (o *WorkflowMoReferenceDataTypeAllOf) SetProperties(v []WorkflowMoReferenceProperty) {
	o.Properties = v
}

func (o WorkflowMoReferenceDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMoReferenceDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMoReferenceDataTypeAllOf := _WorkflowMoReferenceDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowMoReferenceDataTypeAllOf); err == nil {
		*o = WorkflowMoReferenceDataTypeAllOf(varWorkflowMoReferenceDataTypeAllOf)
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

type NullableWorkflowMoReferenceDataTypeAllOf struct {
	value *WorkflowMoReferenceDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowMoReferenceDataTypeAllOf) Get() *WorkflowMoReferenceDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowMoReferenceDataTypeAllOf) Set(val *WorkflowMoReferenceDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoReferenceDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoReferenceDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoReferenceDataTypeAllOf(val *WorkflowMoReferenceDataTypeAllOf) *NullableWorkflowMoReferenceDataTypeAllOf {
	return &NullableWorkflowMoReferenceDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowMoReferenceDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoReferenceDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
