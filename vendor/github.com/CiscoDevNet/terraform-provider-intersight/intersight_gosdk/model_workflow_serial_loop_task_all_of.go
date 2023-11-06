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
)

// WorkflowSerialLoopTaskAllOf Definition of the list of properties defined in 'workflow.SerialLoopTask', excluding properties defined in parent classes.
type WorkflowSerialLoopTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Condition expression which will be evaluated and when expression is true the tasks under loop will be executed.
	Condition            *string `json:"Condition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSerialLoopTaskAllOf WorkflowSerialLoopTaskAllOf

// NewWorkflowSerialLoopTaskAllOf instantiates a new WorkflowSerialLoopTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSerialLoopTaskAllOf(classId string, objectType string) *WorkflowSerialLoopTaskAllOf {
	this := WorkflowSerialLoopTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowSerialLoopTaskAllOfWithDefaults instantiates a new WorkflowSerialLoopTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSerialLoopTaskAllOfWithDefaults() *WorkflowSerialLoopTaskAllOf {
	this := WorkflowSerialLoopTaskAllOf{}
	var classId string = "workflow.SerialLoopTask"
	this.ClassId = classId
	var objectType string = "workflow.SerialLoopTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSerialLoopTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSerialLoopTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSerialLoopTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSerialLoopTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSerialLoopTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSerialLoopTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WorkflowSerialLoopTaskAllOf) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSerialLoopTaskAllOf) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WorkflowSerialLoopTaskAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WorkflowSerialLoopTaskAllOf) SetCondition(v string) {
	o.Condition = &v
}

func (o WorkflowSerialLoopTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSerialLoopTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSerialLoopTaskAllOf := _WorkflowSerialLoopTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSerialLoopTaskAllOf); err == nil {
		*o = WorkflowSerialLoopTaskAllOf(varWorkflowSerialLoopTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSerialLoopTaskAllOf struct {
	value *WorkflowSerialLoopTaskAllOf
	isSet bool
}

func (v NullableWorkflowSerialLoopTaskAllOf) Get() *WorkflowSerialLoopTaskAllOf {
	return v.value
}

func (v *NullableWorkflowSerialLoopTaskAllOf) Set(val *WorkflowSerialLoopTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSerialLoopTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSerialLoopTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSerialLoopTaskAllOf(val *WorkflowSerialLoopTaskAllOf) *NullableWorkflowSerialLoopTaskAllOf {
	return &NullableWorkflowSerialLoopTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSerialLoopTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSerialLoopTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
