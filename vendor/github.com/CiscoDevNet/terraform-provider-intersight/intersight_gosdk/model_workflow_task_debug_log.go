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

// WorkflowTaskDebugLog Debug information captured during task execution, when the enable debug flag is set at the workflow definition level.
type WorkflowTaskDebugLog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A counter for number of retries.
	RetryCount *int64 `json:"RetryCount,omitempty"`
	// Holds information helpful in isolating task failures.
	TaskDebugLogEntries interface{} `json:"TaskDebugLogEntries,omitempty"`
	// The unique identifier for task instance.
	TaskInstId           *string                           `json:"TaskInstId,omitempty"`
	TaskInfo             *WorkflowTaskInfoRelationship     `json:"TaskInfo,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskDebugLog WorkflowTaskDebugLog

// NewWorkflowTaskDebugLog instantiates a new WorkflowTaskDebugLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskDebugLog(classId string, objectType string) *WorkflowTaskDebugLog {
	this := WorkflowTaskDebugLog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskDebugLogWithDefaults instantiates a new WorkflowTaskDebugLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskDebugLogWithDefaults() *WorkflowTaskDebugLog {
	this := WorkflowTaskDebugLog{}
	var classId string = "workflow.TaskDebugLog"
	this.ClassId = classId
	var objectType string = "workflow.TaskDebugLog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskDebugLog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskDebugLog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskDebugLog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskDebugLog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *WorkflowTaskDebugLog) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *WorkflowTaskDebugLog) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *WorkflowTaskDebugLog) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetTaskDebugLogEntries returns the TaskDebugLogEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskDebugLog) GetTaskDebugLogEntries() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaskDebugLogEntries
}

// GetTaskDebugLogEntriesOk returns a tuple with the TaskDebugLogEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskDebugLog) GetTaskDebugLogEntriesOk() (*interface{}, bool) {
	if o == nil || o.TaskDebugLogEntries == nil {
		return nil, false
	}
	return &o.TaskDebugLogEntries, true
}

// HasTaskDebugLogEntries returns a boolean if a field has been set.
func (o *WorkflowTaskDebugLog) HasTaskDebugLogEntries() bool {
	if o != nil && o.TaskDebugLogEntries != nil {
		return true
	}

	return false
}

// SetTaskDebugLogEntries gets a reference to the given interface{} and assigns it to the TaskDebugLogEntries field.
func (o *WorkflowTaskDebugLog) SetTaskDebugLogEntries(v interface{}) {
	o.TaskDebugLogEntries = v
}

// GetTaskInstId returns the TaskInstId field value if set, zero value otherwise.
func (o *WorkflowTaskDebugLog) GetTaskInstId() string {
	if o == nil || o.TaskInstId == nil {
		var ret string
		return ret
	}
	return *o.TaskInstId
}

// GetTaskInstIdOk returns a tuple with the TaskInstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetTaskInstIdOk() (*string, bool) {
	if o == nil || o.TaskInstId == nil {
		return nil, false
	}
	return o.TaskInstId, true
}

// HasTaskInstId returns a boolean if a field has been set.
func (o *WorkflowTaskDebugLog) HasTaskInstId() bool {
	if o != nil && o.TaskInstId != nil {
		return true
	}

	return false
}

// SetTaskInstId gets a reference to the given string and assigns it to the TaskInstId field.
func (o *WorkflowTaskDebugLog) SetTaskInstId(v string) {
	o.TaskInstId = &v
}

// GetTaskInfo returns the TaskInfo field value if set, zero value otherwise.
func (o *WorkflowTaskDebugLog) GetTaskInfo() WorkflowTaskInfoRelationship {
	if o == nil || o.TaskInfo == nil {
		var ret WorkflowTaskInfoRelationship
		return ret
	}
	return *o.TaskInfo
}

// GetTaskInfoOk returns a tuple with the TaskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetTaskInfoOk() (*WorkflowTaskInfoRelationship, bool) {
	if o == nil || o.TaskInfo == nil {
		return nil, false
	}
	return o.TaskInfo, true
}

// HasTaskInfo returns a boolean if a field has been set.
func (o *WorkflowTaskDebugLog) HasTaskInfo() bool {
	if o != nil && o.TaskInfo != nil {
		return true
	}

	return false
}

// SetTaskInfo gets a reference to the given WorkflowTaskInfoRelationship and assigns it to the TaskInfo field.
func (o *WorkflowTaskDebugLog) SetTaskInfo(v WorkflowTaskInfoRelationship) {
	o.TaskInfo = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowTaskDebugLog) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskDebugLog) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowTaskDebugLog) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowTaskDebugLog) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowTaskDebugLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RetryCount != nil {
		toSerialize["RetryCount"] = o.RetryCount
	}
	if o.TaskDebugLogEntries != nil {
		toSerialize["TaskDebugLogEntries"] = o.TaskDebugLogEntries
	}
	if o.TaskInstId != nil {
		toSerialize["TaskInstId"] = o.TaskInstId
	}
	if o.TaskInfo != nil {
		toSerialize["TaskInfo"] = o.TaskInfo
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskDebugLog) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTaskDebugLogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A counter for number of retries.
		RetryCount *int64 `json:"RetryCount,omitempty"`
		// Holds information helpful in isolating task failures.
		TaskDebugLogEntries interface{} `json:"TaskDebugLogEntries,omitempty"`
		// The unique identifier for task instance.
		TaskInstId   *string                           `json:"TaskInstId,omitempty"`
		TaskInfo     *WorkflowTaskInfoRelationship     `json:"TaskInfo,omitempty"`
		WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	}

	varWorkflowTaskDebugLogWithoutEmbeddedStruct := WorkflowTaskDebugLogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTaskDebugLogWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskDebugLog := _WorkflowTaskDebugLog{}
		varWorkflowTaskDebugLog.ClassId = varWorkflowTaskDebugLogWithoutEmbeddedStruct.ClassId
		varWorkflowTaskDebugLog.ObjectType = varWorkflowTaskDebugLogWithoutEmbeddedStruct.ObjectType
		varWorkflowTaskDebugLog.RetryCount = varWorkflowTaskDebugLogWithoutEmbeddedStruct.RetryCount
		varWorkflowTaskDebugLog.TaskDebugLogEntries = varWorkflowTaskDebugLogWithoutEmbeddedStruct.TaskDebugLogEntries
		varWorkflowTaskDebugLog.TaskInstId = varWorkflowTaskDebugLogWithoutEmbeddedStruct.TaskInstId
		varWorkflowTaskDebugLog.TaskInfo = varWorkflowTaskDebugLogWithoutEmbeddedStruct.TaskInfo
		varWorkflowTaskDebugLog.WorkflowInfo = varWorkflowTaskDebugLogWithoutEmbeddedStruct.WorkflowInfo
		*o = WorkflowTaskDebugLog(varWorkflowTaskDebugLog)
	} else {
		return err
	}

	varWorkflowTaskDebugLog := _WorkflowTaskDebugLog{}

	err = json.Unmarshal(bytes, &varWorkflowTaskDebugLog)
	if err == nil {
		o.MoBaseMo = varWorkflowTaskDebugLog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RetryCount")
		delete(additionalProperties, "TaskDebugLogEntries")
		delete(additionalProperties, "TaskInstId")
		delete(additionalProperties, "TaskInfo")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableWorkflowTaskDebugLog struct {
	value *WorkflowTaskDebugLog
	isSet bool
}

func (v NullableWorkflowTaskDebugLog) Get() *WorkflowTaskDebugLog {
	return v.value
}

func (v *NullableWorkflowTaskDebugLog) Set(val *WorkflowTaskDebugLog) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskDebugLog) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskDebugLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskDebugLog(val *WorkflowTaskDebugLog) *NullableWorkflowTaskDebugLog {
	return &NullableWorkflowTaskDebugLog{value: val, isSet: true}
}

func (v NullableWorkflowTaskDebugLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskDebugLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
