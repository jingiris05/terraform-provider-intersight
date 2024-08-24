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

// checks if the IaasMostRunTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasMostRunTasks{}

// IaasMostRunTasks Describes most run workflow tasks within UCSD.
type IaasMostRunTasks struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A functional area to which a task belongs to.
	TaskCategory *string `json:"TaskCategory,omitempty"`
	// Number of times this task has executed.
	TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
	// Name of the task executed in UCSD.
	TaskName *string `json:"TaskName,omitempty"`
	// Type of the task whether it is system task or custom task.
	TaskType             *string                          `json:"TaskType,omitempty"`
	Guid                 NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasMostRunTasks IaasMostRunTasks

// NewIaasMostRunTasks instantiates a new IaasMostRunTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasMostRunTasks(classId string, objectType string) *IaasMostRunTasks {
	this := IaasMostRunTasks{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasMostRunTasksWithDefaults instantiates a new IaasMostRunTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasMostRunTasksWithDefaults() *IaasMostRunTasks {
	this := IaasMostRunTasks{}
	var classId string = "iaas.MostRunTasks"
	this.ClassId = classId
	var objectType string = "iaas.MostRunTasks"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasMostRunTasks) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasMostRunTasks) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.MostRunTasks" of the ClassId field.
func (o *IaasMostRunTasks) GetDefaultClassId() interface{} {
	return "iaas.MostRunTasks"
}

// GetObjectType returns the ObjectType field value
func (o *IaasMostRunTasks) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasMostRunTasks) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.MostRunTasks" of the ObjectType field.
func (o *IaasMostRunTasks) GetDefaultObjectType() interface{} {
	return "iaas.MostRunTasks"
}

// GetTaskCategory returns the TaskCategory field value if set, zero value otherwise.
func (o *IaasMostRunTasks) GetTaskCategory() string {
	if o == nil || IsNil(o.TaskCategory) {
		var ret string
		return ret
	}
	return *o.TaskCategory
}

// GetTaskCategoryOk returns a tuple with the TaskCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetTaskCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.TaskCategory) {
		return nil, false
	}
	return o.TaskCategory, true
}

// HasTaskCategory returns a boolean if a field has been set.
func (o *IaasMostRunTasks) HasTaskCategory() bool {
	if o != nil && !IsNil(o.TaskCategory) {
		return true
	}

	return false
}

// SetTaskCategory gets a reference to the given string and assigns it to the TaskCategory field.
func (o *IaasMostRunTasks) SetTaskCategory(v string) {
	o.TaskCategory = &v
}

// GetTaskExecutionCount returns the TaskExecutionCount field value if set, zero value otherwise.
func (o *IaasMostRunTasks) GetTaskExecutionCount() int64 {
	if o == nil || IsNil(o.TaskExecutionCount) {
		var ret int64
		return ret
	}
	return *o.TaskExecutionCount
}

// GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetTaskExecutionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TaskExecutionCount) {
		return nil, false
	}
	return o.TaskExecutionCount, true
}

// HasTaskExecutionCount returns a boolean if a field has been set.
func (o *IaasMostRunTasks) HasTaskExecutionCount() bool {
	if o != nil && !IsNil(o.TaskExecutionCount) {
		return true
	}

	return false
}

// SetTaskExecutionCount gets a reference to the given int64 and assigns it to the TaskExecutionCount field.
func (o *IaasMostRunTasks) SetTaskExecutionCount(v int64) {
	o.TaskExecutionCount = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *IaasMostRunTasks) GetTaskName() string {
	if o == nil || IsNil(o.TaskName) {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetTaskNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaskName) {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *IaasMostRunTasks) HasTaskName() bool {
	if o != nil && !IsNil(o.TaskName) {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *IaasMostRunTasks) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *IaasMostRunTasks) GetTaskType() string {
	if o == nil || IsNil(o.TaskType) {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasks) GetTaskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *IaasMostRunTasks) HasTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *IaasMostRunTasks) SetTaskType(v string) {
	o.TaskType = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasMostRunTasks) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasMostRunTasks) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasMostRunTasks) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableIaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasMostRunTasks) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid.Set(&v)
}

// SetGuidNil sets the value for Guid to be an explicit nil
func (o *IaasMostRunTasks) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *IaasMostRunTasks) UnsetGuid() {
	o.Guid.Unset()
}

func (o IaasMostRunTasks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasMostRunTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.TaskCategory) {
		toSerialize["TaskCategory"] = o.TaskCategory
	}
	if !IsNil(o.TaskExecutionCount) {
		toSerialize["TaskExecutionCount"] = o.TaskExecutionCount
	}
	if !IsNil(o.TaskName) {
		toSerialize["TaskName"] = o.TaskName
	}
	if !IsNil(o.TaskType) {
		toSerialize["TaskType"] = o.TaskType
	}
	if o.Guid.IsSet() {
		toSerialize["Guid"] = o.Guid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasMostRunTasks) UnmarshalJSON(data []byte) (err error) {
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
	type IaasMostRunTasksWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A functional area to which a task belongs to.
		TaskCategory *string `json:"TaskCategory,omitempty"`
		// Number of times this task has executed.
		TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
		// Name of the task executed in UCSD.
		TaskName *string `json:"TaskName,omitempty"`
		// Type of the task whether it is system task or custom task.
		TaskType *string                          `json:"TaskType,omitempty"`
		Guid     NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasMostRunTasksWithoutEmbeddedStruct := IaasMostRunTasksWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasMostRunTasksWithoutEmbeddedStruct)
	if err == nil {
		varIaasMostRunTasks := _IaasMostRunTasks{}
		varIaasMostRunTasks.ClassId = varIaasMostRunTasksWithoutEmbeddedStruct.ClassId
		varIaasMostRunTasks.ObjectType = varIaasMostRunTasksWithoutEmbeddedStruct.ObjectType
		varIaasMostRunTasks.TaskCategory = varIaasMostRunTasksWithoutEmbeddedStruct.TaskCategory
		varIaasMostRunTasks.TaskExecutionCount = varIaasMostRunTasksWithoutEmbeddedStruct.TaskExecutionCount
		varIaasMostRunTasks.TaskName = varIaasMostRunTasksWithoutEmbeddedStruct.TaskName
		varIaasMostRunTasks.TaskType = varIaasMostRunTasksWithoutEmbeddedStruct.TaskType
		varIaasMostRunTasks.Guid = varIaasMostRunTasksWithoutEmbeddedStruct.Guid
		*o = IaasMostRunTasks(varIaasMostRunTasks)
	} else {
		return err
	}

	varIaasMostRunTasks := _IaasMostRunTasks{}

	err = json.Unmarshal(data, &varIaasMostRunTasks)
	if err == nil {
		o.MoBaseMo = varIaasMostRunTasks.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TaskCategory")
		delete(additionalProperties, "TaskExecutionCount")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "TaskType")
		delete(additionalProperties, "Guid")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIaasMostRunTasks struct {
	value *IaasMostRunTasks
	isSet bool
}

func (v NullableIaasMostRunTasks) Get() *IaasMostRunTasks {
	return v.value
}

func (v *NullableIaasMostRunTasks) Set(val *IaasMostRunTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasMostRunTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasMostRunTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasMostRunTasks(val *IaasMostRunTasks) *NullableIaasMostRunTasks {
	return &NullableIaasMostRunTasks{value: val, isSet: true}
}

func (v NullableIaasMostRunTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasMostRunTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
