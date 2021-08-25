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

// WorkflowRollbackTaskAllOf Definition of the list of properties defined in 'workflow.RollbackTask', excluding properties defined in parent classes.
type WorkflowRollbackTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The catalog under which the task definition has been added.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// Description of rollback task definition.
	Description *string `json:"Description,omitempty"`
	// Input parameters mapping for rollback task from the input or output of the main task definition.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// Name of the task definition which is capable of doing rollback of this task.
	Name *string `json:"Name,omitempty"`
	// The resolved referenced rollback task definition managed object.
	TaskMoid *string `json:"TaskMoid,omitempty"`
	// The version of the task definition.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowRollbackTaskAllOf WorkflowRollbackTaskAllOf

// NewWorkflowRollbackTaskAllOf instantiates a new WorkflowRollbackTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRollbackTaskAllOf(classId string, objectType string) *WorkflowRollbackTaskAllOf {
	this := WorkflowRollbackTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowRollbackTaskAllOfWithDefaults instantiates a new WorkflowRollbackTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRollbackTaskAllOfWithDefaults() *WorkflowRollbackTaskAllOf {
	this := WorkflowRollbackTaskAllOf{}
	var classId string = "workflow.RollbackTask"
	this.ClassId = classId
	var objectType string = "workflow.RollbackTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowRollbackTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowRollbackTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowRollbackTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowRollbackTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackTaskAllOf) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowRollbackTaskAllOf) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowRollbackTaskAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowRollbackTaskAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRollbackTaskAllOf) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRollbackTaskAllOf) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowRollbackTaskAllOf) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRollbackTaskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRollbackTaskAllOf) SetName(v string) {
	o.Name = &v
}

// GetTaskMoid returns the TaskMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackTaskAllOf) GetTaskMoid() string {
	if o == nil || o.TaskMoid == nil {
		var ret string
		return ret
	}
	return *o.TaskMoid
}

// GetTaskMoidOk returns a tuple with the TaskMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetTaskMoidOk() (*string, bool) {
	if o == nil || o.TaskMoid == nil {
		return nil, false
	}
	return o.TaskMoid, true
}

// HasTaskMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasTaskMoid() bool {
	if o != nil && o.TaskMoid != nil {
		return true
	}

	return false
}

// SetTaskMoid gets a reference to the given string and assigns it to the TaskMoid field.
func (o *WorkflowRollbackTaskAllOf) SetTaskMoid(v string) {
	o.TaskMoid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowRollbackTaskAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTaskAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowRollbackTaskAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowRollbackTaskAllOf) SetVersion(v int64) {
	o.Version = &v
}

func (o WorkflowRollbackTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TaskMoid != nil {
		toSerialize["TaskMoid"] = o.TaskMoid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowRollbackTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowRollbackTaskAllOf := _WorkflowRollbackTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowRollbackTaskAllOf); err == nil {
		*o = WorkflowRollbackTaskAllOf(varWorkflowRollbackTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TaskMoid")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowRollbackTaskAllOf struct {
	value *WorkflowRollbackTaskAllOf
	isSet bool
}

func (v NullableWorkflowRollbackTaskAllOf) Get() *WorkflowRollbackTaskAllOf {
	return v.value
}

func (v *NullableWorkflowRollbackTaskAllOf) Set(val *WorkflowRollbackTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRollbackTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRollbackTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRollbackTaskAllOf(val *WorkflowRollbackTaskAllOf) *NullableWorkflowRollbackTaskAllOf {
	return &NullableWorkflowRollbackTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowRollbackTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRollbackTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
