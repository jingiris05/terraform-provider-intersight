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

// checks if the WorkflowRollbackTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRollbackTask{}

// WorkflowRollbackTask Rollback task mapping information.
type WorkflowRollbackTask struct {
	MoBaseComplexType
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
	// The rollback task will not be executed if the given condition evaluates to \"true\".
	SkipCondition *string `json:"SkipCondition,omitempty"`
	// The resolved referenced rollback task definition managed object.
	TaskMoid *string `json:"TaskMoid,omitempty"`
	// The version of the task definition.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowRollbackTask WorkflowRollbackTask

// NewWorkflowRollbackTask instantiates a new WorkflowRollbackTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRollbackTask(classId string, objectType string) *WorkflowRollbackTask {
	this := WorkflowRollbackTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	var version int64 = 1
	this.Version = &version
	return &this
}

// NewWorkflowRollbackTaskWithDefaults instantiates a new WorkflowRollbackTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRollbackTaskWithDefaults() *WorkflowRollbackTask {
	this := WorkflowRollbackTask{}
	var classId string = "workflow.RollbackTask"
	this.ClassId = classId
	var objectType string = "workflow.RollbackTask"
	this.ObjectType = objectType
	var version int64 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowRollbackTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowRollbackTask) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.RollbackTask" of the ClassId field.
func (o *WorkflowRollbackTask) GetDefaultClassId() interface{} {
	return "workflow.RollbackTask"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowRollbackTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowRollbackTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.RollbackTask" of the ObjectType field.
func (o *WorkflowRollbackTask) GetDefaultObjectType() interface{} {
	return "workflow.RollbackTask"
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetCatalogMoid() string {
	if o == nil || IsNil(o.CatalogMoid) {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetCatalogMoidOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogMoid) {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasCatalogMoid() bool {
	if o != nil && !IsNil(o.CatalogMoid) {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowRollbackTask) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowRollbackTask) SetDescription(v string) {
	o.Description = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRollbackTask) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRollbackTask) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InputParameters) {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasInputParameters() bool {
	if o != nil && !IsNil(o.InputParameters) {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowRollbackTask) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRollbackTask) SetName(v string) {
	o.Name = &v
}

// GetSkipCondition returns the SkipCondition field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetSkipCondition() string {
	if o == nil || IsNil(o.SkipCondition) {
		var ret string
		return ret
	}
	return *o.SkipCondition
}

// GetSkipConditionOk returns a tuple with the SkipCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetSkipConditionOk() (*string, bool) {
	if o == nil || IsNil(o.SkipCondition) {
		return nil, false
	}
	return o.SkipCondition, true
}

// HasSkipCondition returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasSkipCondition() bool {
	if o != nil && !IsNil(o.SkipCondition) {
		return true
	}

	return false
}

// SetSkipCondition gets a reference to the given string and assigns it to the SkipCondition field.
func (o *WorkflowRollbackTask) SetSkipCondition(v string) {
	o.SkipCondition = &v
}

// GetTaskMoid returns the TaskMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetTaskMoid() string {
	if o == nil || IsNil(o.TaskMoid) {
		var ret string
		return ret
	}
	return *o.TaskMoid
}

// GetTaskMoidOk returns a tuple with the TaskMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetTaskMoidOk() (*string, bool) {
	if o == nil || IsNil(o.TaskMoid) {
		return nil, false
	}
	return o.TaskMoid, true
}

// HasTaskMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasTaskMoid() bool {
	if o != nil && !IsNil(o.TaskMoid) {
		return true
	}

	return false
}

// SetTaskMoid gets a reference to the given string and assigns it to the TaskMoid field.
func (o *WorkflowRollbackTask) SetTaskMoid(v string) {
	o.TaskMoid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowRollbackTask) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackTask) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowRollbackTask) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowRollbackTask) SetVersion(v int64) {
	o.Version = &v
}

func (o WorkflowRollbackTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRollbackTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CatalogMoid) {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SkipCondition) {
		toSerialize["SkipCondition"] = o.SkipCondition
	}
	if !IsNil(o.TaskMoid) {
		toSerialize["TaskMoid"] = o.TaskMoid
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowRollbackTask) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowRollbackTaskWithoutEmbeddedStruct struct {
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
		// The rollback task will not be executed if the given condition evaluates to \"true\".
		SkipCondition *string `json:"SkipCondition,omitempty"`
		// The resolved referenced rollback task definition managed object.
		TaskMoid *string `json:"TaskMoid,omitempty"`
		// The version of the task definition.
		Version *int64 `json:"Version,omitempty"`
	}

	varWorkflowRollbackTaskWithoutEmbeddedStruct := WorkflowRollbackTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowRollbackTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowRollbackTask := _WorkflowRollbackTask{}
		varWorkflowRollbackTask.ClassId = varWorkflowRollbackTaskWithoutEmbeddedStruct.ClassId
		varWorkflowRollbackTask.ObjectType = varWorkflowRollbackTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowRollbackTask.CatalogMoid = varWorkflowRollbackTaskWithoutEmbeddedStruct.CatalogMoid
		varWorkflowRollbackTask.Description = varWorkflowRollbackTaskWithoutEmbeddedStruct.Description
		varWorkflowRollbackTask.InputParameters = varWorkflowRollbackTaskWithoutEmbeddedStruct.InputParameters
		varWorkflowRollbackTask.Name = varWorkflowRollbackTaskWithoutEmbeddedStruct.Name
		varWorkflowRollbackTask.SkipCondition = varWorkflowRollbackTaskWithoutEmbeddedStruct.SkipCondition
		varWorkflowRollbackTask.TaskMoid = varWorkflowRollbackTaskWithoutEmbeddedStruct.TaskMoid
		varWorkflowRollbackTask.Version = varWorkflowRollbackTaskWithoutEmbeddedStruct.Version
		*o = WorkflowRollbackTask(varWorkflowRollbackTask)
	} else {
		return err
	}

	varWorkflowRollbackTask := _WorkflowRollbackTask{}

	err = json.Unmarshal(data, &varWorkflowRollbackTask)
	if err == nil {
		o.MoBaseComplexType = varWorkflowRollbackTask.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SkipCondition")
		delete(additionalProperties, "TaskMoid")
		delete(additionalProperties, "Version")

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

type NullableWorkflowRollbackTask struct {
	value *WorkflowRollbackTask
	isSet bool
}

func (v NullableWorkflowRollbackTask) Get() *WorkflowRollbackTask {
	return v.value
}

func (v *NullableWorkflowRollbackTask) Set(val *WorkflowRollbackTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRollbackTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRollbackTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRollbackTask(val *WorkflowRollbackTask) *NullableWorkflowRollbackTask {
	return &NullableWorkflowRollbackTask{value: val, isSet: true}
}

func (v NullableWorkflowRollbackTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRollbackTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
