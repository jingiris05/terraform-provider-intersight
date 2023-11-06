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

// WorkflowTaskRetryInfoAllOf Definition of the list of properties defined in 'workflow.TaskRetryInfo', excluding properties defined in parent classes.
type WorkflowTaskRetryInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the retried task.
	Status *string `json:"Status,omitempty"`
	// Retry instance will get a unique instance id.
	TaskInstId           *string `json:"TaskInstId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskRetryInfoAllOf WorkflowTaskRetryInfoAllOf

// NewWorkflowTaskRetryInfoAllOf instantiates a new WorkflowTaskRetryInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskRetryInfoAllOf(classId string, objectType string) *WorkflowTaskRetryInfoAllOf {
	this := WorkflowTaskRetryInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskRetryInfoAllOfWithDefaults instantiates a new WorkflowTaskRetryInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskRetryInfoAllOfWithDefaults() *WorkflowTaskRetryInfoAllOf {
	this := WorkflowTaskRetryInfoAllOf{}
	var classId string = "workflow.TaskRetryInfo"
	this.ClassId = classId
	var objectType string = "workflow.TaskRetryInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskRetryInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskRetryInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskRetryInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskRetryInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskRetryInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskRetryInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowTaskRetryInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskRetryInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowTaskRetryInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowTaskRetryInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskInstId returns the TaskInstId field value if set, zero value otherwise.
func (o *WorkflowTaskRetryInfoAllOf) GetTaskInstId() string {
	if o == nil || o.TaskInstId == nil {
		var ret string
		return ret
	}
	return *o.TaskInstId
}

// GetTaskInstIdOk returns a tuple with the TaskInstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskRetryInfoAllOf) GetTaskInstIdOk() (*string, bool) {
	if o == nil || o.TaskInstId == nil {
		return nil, false
	}
	return o.TaskInstId, true
}

// HasTaskInstId returns a boolean if a field has been set.
func (o *WorkflowTaskRetryInfoAllOf) HasTaskInstId() bool {
	if o != nil && o.TaskInstId != nil {
		return true
	}

	return false
}

// SetTaskInstId gets a reference to the given string and assigns it to the TaskInstId field.
func (o *WorkflowTaskRetryInfoAllOf) SetTaskInstId(v string) {
	o.TaskInstId = &v
}

func (o WorkflowTaskRetryInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskInstId != nil {
		toSerialize["TaskInstId"] = o.TaskInstId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskRetryInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTaskRetryInfoAllOf := _WorkflowTaskRetryInfoAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTaskRetryInfoAllOf); err == nil {
		*o = WorkflowTaskRetryInfoAllOf(varWorkflowTaskRetryInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskInstId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskRetryInfoAllOf struct {
	value *WorkflowTaskRetryInfoAllOf
	isSet bool
}

func (v NullableWorkflowTaskRetryInfoAllOf) Get() *WorkflowTaskRetryInfoAllOf {
	return v.value
}

func (v *NullableWorkflowTaskRetryInfoAllOf) Set(val *WorkflowTaskRetryInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskRetryInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskRetryInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskRetryInfoAllOf(val *WorkflowTaskRetryInfoAllOf) *NullableWorkflowTaskRetryInfoAllOf {
	return &NullableWorkflowTaskRetryInfoAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskRetryInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskRetryInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
