/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14237
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowCustomDataType This data type represents a custom data object.
type WorkflowCustomDataType struct {
	WorkflowBaseDataType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Properties           NullableWorkflowCustomDataProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomDataType WorkflowCustomDataType

// NewWorkflowCustomDataType instantiates a new WorkflowCustomDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataType(classId string, objectType string) *WorkflowCustomDataType {
	this := WorkflowCustomDataType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCustomDataTypeWithDefaults instantiates a new WorkflowCustomDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypeWithDefaults() *WorkflowCustomDataType {
	this := WorkflowCustomDataType{}
	var classId string = "workflow.CustomDataType"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataType) GetProperties() WorkflowCustomDataProperty {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowCustomDataProperty
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataType) GetPropertiesOk() (*WorkflowCustomDataProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowCustomDataType) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowCustomDataProperty and assigns it to the Properties field.
func (o *WorkflowCustomDataType) SetProperties(v WorkflowCustomDataProperty) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowCustomDataType) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowCustomDataType) UnsetProperties() {
	o.Properties.Unset()
}

func (o WorkflowCustomDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseDataType, errWorkflowBaseDataType := json.Marshal(o.WorkflowBaseDataType)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
	errWorkflowBaseDataType = json.Unmarshal([]byte(serializedWorkflowBaseDataType), &toSerialize)
	if errWorkflowBaseDataType != nil {
		return []byte{}, errWorkflowBaseDataType
	}
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

func (o *WorkflowCustomDataType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCustomDataTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                             `json:"ObjectType"`
		Properties NullableWorkflowCustomDataProperty `json:"Properties,omitempty"`
	}

	varWorkflowCustomDataTypeWithoutEmbeddedStruct := WorkflowCustomDataTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCustomDataType := _WorkflowCustomDataType{}
		varWorkflowCustomDataType.ClassId = varWorkflowCustomDataTypeWithoutEmbeddedStruct.ClassId
		varWorkflowCustomDataType.ObjectType = varWorkflowCustomDataTypeWithoutEmbeddedStruct.ObjectType
		varWorkflowCustomDataType.Properties = varWorkflowCustomDataTypeWithoutEmbeddedStruct.Properties
		*o = WorkflowCustomDataType(varWorkflowCustomDataType)
	} else {
		return err
	}

	varWorkflowCustomDataType := _WorkflowCustomDataType{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataType)
	if err == nil {
		o.WorkflowBaseDataType = varWorkflowCustomDataType.WorkflowBaseDataType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Properties")

		// remove fields from embedded structs
		reflectWorkflowBaseDataType := reflect.ValueOf(o.WorkflowBaseDataType)
		for i := 0; i < reflectWorkflowBaseDataType.Type().NumField(); i++ {
			t := reflectWorkflowBaseDataType.Type().Field(i)

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

type NullableWorkflowCustomDataType struct {
	value *WorkflowCustomDataType
	isSet bool
}

func (v NullableWorkflowCustomDataType) Get() *WorkflowCustomDataType {
	return v.value
}

func (v *NullableWorkflowCustomDataType) Set(val *WorkflowCustomDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataType(val *WorkflowCustomDataType) *NullableWorkflowCustomDataType {
	return &NullableWorkflowCustomDataType{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
