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
	"time"
)

// checks if the SchedulerTaskScheduleStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerTaskScheduleStatus{}

// SchedulerTaskScheduleStatus The scheduled task details.
type SchedulerTaskScheduleStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The task completion count, which includes both successful executions and any failures.
	Count *int64 `json:"Count,omitempty"`
	// The status of the current task. * `None` - No status is set (default). * `Scheduled` - The status is set when a task is scheduled. * `Running` - The status is set when a task is running. * `Completed` - The status is set when a task is complete. * `Failed` - The status is set when a task fails. * `Suspended` - The status is set when a task is suspended. * `Skipped` - The status is set when a task is skipped because the previous task is still running.
	CurrentStatus *string `json:"CurrentStatus,omitempty"`
	// Indicates if this task was suspended by the system.
	IsSystemSuspended *bool `json:"IsSystemSuspended,omitempty"`
	// The next run time for a recurrently scheduled the task.
	NextRunStartTime *time.Time `json:"NextRunStartTime,omitempty"`
	// The time when the last occurrence of scheduled task completed.
	PrevRunEndTime *time.Time `json:"PrevRunEndTime,omitempty"`
	// The previous time the scheduled task was run.
	PrevRunStartTime *time.Time `json:"PrevRunStartTime,omitempty"`
	// The reason why the task failed or suspended.
	Reason               *string `json:"Reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerTaskScheduleStatus SchedulerTaskScheduleStatus

// NewSchedulerTaskScheduleStatus instantiates a new SchedulerTaskScheduleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerTaskScheduleStatus(classId string, objectType string) *SchedulerTaskScheduleStatus {
	this := SchedulerTaskScheduleStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSchedulerTaskScheduleStatusWithDefaults instantiates a new SchedulerTaskScheduleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerTaskScheduleStatusWithDefaults() *SchedulerTaskScheduleStatus {
	this := SchedulerTaskScheduleStatus{}
	var classId string = "scheduler.TaskScheduleStatus"
	this.ClassId = classId
	var objectType string = "scheduler.TaskScheduleStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerTaskScheduleStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerTaskScheduleStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.TaskScheduleStatus" of the ClassId field.
func (o *SchedulerTaskScheduleStatus) GetDefaultClassId() interface{} {
	return "scheduler.TaskScheduleStatus"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerTaskScheduleStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerTaskScheduleStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.TaskScheduleStatus" of the ObjectType field.
func (o *SchedulerTaskScheduleStatus) GetDefaultObjectType() interface{} {
	return "scheduler.TaskScheduleStatus"
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SchedulerTaskScheduleStatus) SetCount(v int64) {
	o.Count = &v
}

// GetCurrentStatus returns the CurrentStatus field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetCurrentStatus() string {
	if o == nil || IsNil(o.CurrentStatus) {
		var ret string
		return ret
	}
	return *o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetCurrentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentStatus) {
		return nil, false
	}
	return o.CurrentStatus, true
}

// HasCurrentStatus returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasCurrentStatus() bool {
	if o != nil && !IsNil(o.CurrentStatus) {
		return true
	}

	return false
}

// SetCurrentStatus gets a reference to the given string and assigns it to the CurrentStatus field.
func (o *SchedulerTaskScheduleStatus) SetCurrentStatus(v string) {
	o.CurrentStatus = &v
}

// GetIsSystemSuspended returns the IsSystemSuspended field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetIsSystemSuspended() bool {
	if o == nil || IsNil(o.IsSystemSuspended) {
		var ret bool
		return ret
	}
	return *o.IsSystemSuspended
}

// GetIsSystemSuspendedOk returns a tuple with the IsSystemSuspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetIsSystemSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystemSuspended) {
		return nil, false
	}
	return o.IsSystemSuspended, true
}

// HasIsSystemSuspended returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasIsSystemSuspended() bool {
	if o != nil && !IsNil(o.IsSystemSuspended) {
		return true
	}

	return false
}

// SetIsSystemSuspended gets a reference to the given bool and assigns it to the IsSystemSuspended field.
func (o *SchedulerTaskScheduleStatus) SetIsSystemSuspended(v bool) {
	o.IsSystemSuspended = &v
}

// GetNextRunStartTime returns the NextRunStartTime field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetNextRunStartTime() time.Time {
	if o == nil || IsNil(o.NextRunStartTime) {
		var ret time.Time
		return ret
	}
	return *o.NextRunStartTime
}

// GetNextRunStartTimeOk returns a tuple with the NextRunStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetNextRunStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextRunStartTime) {
		return nil, false
	}
	return o.NextRunStartTime, true
}

// HasNextRunStartTime returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasNextRunStartTime() bool {
	if o != nil && !IsNil(o.NextRunStartTime) {
		return true
	}

	return false
}

// SetNextRunStartTime gets a reference to the given time.Time and assigns it to the NextRunStartTime field.
func (o *SchedulerTaskScheduleStatus) SetNextRunStartTime(v time.Time) {
	o.NextRunStartTime = &v
}

// GetPrevRunEndTime returns the PrevRunEndTime field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetPrevRunEndTime() time.Time {
	if o == nil || IsNil(o.PrevRunEndTime) {
		var ret time.Time
		return ret
	}
	return *o.PrevRunEndTime
}

// GetPrevRunEndTimeOk returns a tuple with the PrevRunEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetPrevRunEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrevRunEndTime) {
		return nil, false
	}
	return o.PrevRunEndTime, true
}

// HasPrevRunEndTime returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasPrevRunEndTime() bool {
	if o != nil && !IsNil(o.PrevRunEndTime) {
		return true
	}

	return false
}

// SetPrevRunEndTime gets a reference to the given time.Time and assigns it to the PrevRunEndTime field.
func (o *SchedulerTaskScheduleStatus) SetPrevRunEndTime(v time.Time) {
	o.PrevRunEndTime = &v
}

// GetPrevRunStartTime returns the PrevRunStartTime field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetPrevRunStartTime() time.Time {
	if o == nil || IsNil(o.PrevRunStartTime) {
		var ret time.Time
		return ret
	}
	return *o.PrevRunStartTime
}

// GetPrevRunStartTimeOk returns a tuple with the PrevRunStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetPrevRunStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrevRunStartTime) {
		return nil, false
	}
	return o.PrevRunStartTime, true
}

// HasPrevRunStartTime returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasPrevRunStartTime() bool {
	if o != nil && !IsNil(o.PrevRunStartTime) {
		return true
	}

	return false
}

// SetPrevRunStartTime gets a reference to the given time.Time and assigns it to the PrevRunStartTime field.
func (o *SchedulerTaskScheduleStatus) SetPrevRunStartTime(v time.Time) {
	o.PrevRunStartTime = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SchedulerTaskScheduleStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerTaskScheduleStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SchedulerTaskScheduleStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SchedulerTaskScheduleStatus) SetReason(v string) {
	o.Reason = &v
}

func (o SchedulerTaskScheduleStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerTaskScheduleStatus) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Count) {
		toSerialize["Count"] = o.Count
	}
	if !IsNil(o.CurrentStatus) {
		toSerialize["CurrentStatus"] = o.CurrentStatus
	}
	if !IsNil(o.IsSystemSuspended) {
		toSerialize["IsSystemSuspended"] = o.IsSystemSuspended
	}
	if !IsNil(o.NextRunStartTime) {
		toSerialize["NextRunStartTime"] = o.NextRunStartTime
	}
	if !IsNil(o.PrevRunEndTime) {
		toSerialize["PrevRunEndTime"] = o.PrevRunEndTime
	}
	if !IsNil(o.PrevRunStartTime) {
		toSerialize["PrevRunStartTime"] = o.PrevRunStartTime
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerTaskScheduleStatus) UnmarshalJSON(data []byte) (err error) {
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
	type SchedulerTaskScheduleStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The task completion count, which includes both successful executions and any failures.
		Count *int64 `json:"Count,omitempty"`
		// The status of the current task. * `None` - No status is set (default). * `Scheduled` - The status is set when a task is scheduled. * `Running` - The status is set when a task is running. * `Completed` - The status is set when a task is complete. * `Failed` - The status is set when a task fails. * `Suspended` - The status is set when a task is suspended. * `Skipped` - The status is set when a task is skipped because the previous task is still running.
		CurrentStatus *string `json:"CurrentStatus,omitempty"`
		// Indicates if this task was suspended by the system.
		IsSystemSuspended *bool `json:"IsSystemSuspended,omitempty"`
		// The next run time for a recurrently scheduled the task.
		NextRunStartTime *time.Time `json:"NextRunStartTime,omitempty"`
		// The time when the last occurrence of scheduled task completed.
		PrevRunEndTime *time.Time `json:"PrevRunEndTime,omitempty"`
		// The previous time the scheduled task was run.
		PrevRunStartTime *time.Time `json:"PrevRunStartTime,omitempty"`
		// The reason why the task failed or suspended.
		Reason *string `json:"Reason,omitempty"`
	}

	varSchedulerTaskScheduleStatusWithoutEmbeddedStruct := SchedulerTaskScheduleStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerTaskScheduleStatusWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerTaskScheduleStatus := _SchedulerTaskScheduleStatus{}
		varSchedulerTaskScheduleStatus.ClassId = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.ClassId
		varSchedulerTaskScheduleStatus.ObjectType = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.ObjectType
		varSchedulerTaskScheduleStatus.Count = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.Count
		varSchedulerTaskScheduleStatus.CurrentStatus = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.CurrentStatus
		varSchedulerTaskScheduleStatus.IsSystemSuspended = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.IsSystemSuspended
		varSchedulerTaskScheduleStatus.NextRunStartTime = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.NextRunStartTime
		varSchedulerTaskScheduleStatus.PrevRunEndTime = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.PrevRunEndTime
		varSchedulerTaskScheduleStatus.PrevRunStartTime = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.PrevRunStartTime
		varSchedulerTaskScheduleStatus.Reason = varSchedulerTaskScheduleStatusWithoutEmbeddedStruct.Reason
		*o = SchedulerTaskScheduleStatus(varSchedulerTaskScheduleStatus)
	} else {
		return err
	}

	varSchedulerTaskScheduleStatus := _SchedulerTaskScheduleStatus{}

	err = json.Unmarshal(data, &varSchedulerTaskScheduleStatus)
	if err == nil {
		o.MoBaseComplexType = varSchedulerTaskScheduleStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "CurrentStatus")
		delete(additionalProperties, "IsSystemSuspended")
		delete(additionalProperties, "NextRunStartTime")
		delete(additionalProperties, "PrevRunEndTime")
		delete(additionalProperties, "PrevRunStartTime")
		delete(additionalProperties, "Reason")

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

type NullableSchedulerTaskScheduleStatus struct {
	value *SchedulerTaskScheduleStatus
	isSet bool
}

func (v NullableSchedulerTaskScheduleStatus) Get() *SchedulerTaskScheduleStatus {
	return v.value
}

func (v *NullableSchedulerTaskScheduleStatus) Set(val *SchedulerTaskScheduleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerTaskScheduleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerTaskScheduleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerTaskScheduleStatus(val *SchedulerTaskScheduleStatus) *NullableSchedulerTaskScheduleStatus {
	return &NullableSchedulerTaskScheduleStatus{value: val, isSet: true}
}

func (v NullableSchedulerTaskScheduleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerTaskScheduleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
