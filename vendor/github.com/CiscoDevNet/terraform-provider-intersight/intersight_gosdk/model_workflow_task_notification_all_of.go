/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowTaskNotificationAllOf Definition of the list of properties defined in 'workflow.TaskNotification', excluding properties defined in parent classes.
type WorkflowTaskNotificationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The correlation id of the scheduled task.
	CorrelationId *string `json:"CorrelationId,omitempty"`
	// The end time of the scheduled task.
	EndTime *string `json:"EndTime,omitempty"`
	// The input of the scheduled task.
	Input *string `json:"Input,omitempty"`
	// The output of the scheduled task.
	Output *string `json:"Output,omitempty"`
	// The reason for incompletion status of the task.
	ReasonForIncompletion *string `json:"ReasonForIncompletion,omitempty"`
	// The task reference name of the scheduled task.
	ReferenceTaskName *string `json:"ReferenceTaskName,omitempty"`
	// The number of times the task retries on failure.
	RetryCount *float32 `json:"RetryCount,omitempty"`
	// The scheduled time of the task.
	ScheduledTime *string `json:"ScheduledTime,omitempty"`
	// The start time of the scheduled task.
	StartTime *string `json:"StartTime,omitempty"`
	// The status of the scheduled task.
	Status *string `json:"Status,omitempty"`
	// The definition of the task explains about the task.
	TaskDefName *string `json:"TaskDefName,omitempty"`
	// The description of the task explains about the task.
	TaskDescription *string `json:"TaskDescription,omitempty"`
	// Unique id of the scheduled task.
	TaskId *string `json:"TaskId,omitempty"`
	// The type of the scheduled task.
	TaskType *string `json:"TaskType,omitempty"`
	// The update time of the scheduled task.
	UpdateTime *string `json:"UpdateTime,omitempty"`
	// The unique id of the running workflow containing this scheduled task.
	WorkflowId *string `json:"WorkflowId,omitempty"`
	// The type of the workflow task.
	WorkflowTaskType *string `json:"WorkflowTaskType,omitempty"`
	// The type of workflow containing this scheduled task.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskNotificationAllOf WorkflowTaskNotificationAllOf

// NewWorkflowTaskNotificationAllOf instantiates a new WorkflowTaskNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskNotificationAllOf(classId string, objectType string) *WorkflowTaskNotificationAllOf {
	this := WorkflowTaskNotificationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskNotificationAllOfWithDefaults instantiates a new WorkflowTaskNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskNotificationAllOfWithDefaults() *WorkflowTaskNotificationAllOf {
	this := WorkflowTaskNotificationAllOf{}
	var classId string = "workflow.TaskNotification"
	this.ClassId = classId
	var objectType string = "workflow.TaskNotification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskNotificationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskNotificationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskNotificationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskNotificationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetCorrelationId() string {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetCorrelationIdOk() (*string, bool) {
	if o == nil || o.CorrelationId == nil {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasCorrelationId() bool {
	if o != nil && o.CorrelationId != nil {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *WorkflowTaskNotificationAllOf) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *WorkflowTaskNotificationAllOf) SetEndTime(v string) {
	o.EndTime = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *WorkflowTaskNotificationAllOf) SetInput(v string) {
	o.Input = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *WorkflowTaskNotificationAllOf) SetOutput(v string) {
	o.Output = &v
}

// GetReasonForIncompletion returns the ReasonForIncompletion field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetReasonForIncompletion() string {
	if o == nil || o.ReasonForIncompletion == nil {
		var ret string
		return ret
	}
	return *o.ReasonForIncompletion
}

// GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetReasonForIncompletionOk() (*string, bool) {
	if o == nil || o.ReasonForIncompletion == nil {
		return nil, false
	}
	return o.ReasonForIncompletion, true
}

// HasReasonForIncompletion returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasReasonForIncompletion() bool {
	if o != nil && o.ReasonForIncompletion != nil {
		return true
	}

	return false
}

// SetReasonForIncompletion gets a reference to the given string and assigns it to the ReasonForIncompletion field.
func (o *WorkflowTaskNotificationAllOf) SetReasonForIncompletion(v string) {
	o.ReasonForIncompletion = &v
}

// GetReferenceTaskName returns the ReferenceTaskName field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetReferenceTaskName() string {
	if o == nil || o.ReferenceTaskName == nil {
		var ret string
		return ret
	}
	return *o.ReferenceTaskName
}

// GetReferenceTaskNameOk returns a tuple with the ReferenceTaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetReferenceTaskNameOk() (*string, bool) {
	if o == nil || o.ReferenceTaskName == nil {
		return nil, false
	}
	return o.ReferenceTaskName, true
}

// HasReferenceTaskName returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasReferenceTaskName() bool {
	if o != nil && o.ReferenceTaskName != nil {
		return true
	}

	return false
}

// SetReferenceTaskName gets a reference to the given string and assigns it to the ReferenceTaskName field.
func (o *WorkflowTaskNotificationAllOf) SetReferenceTaskName(v string) {
	o.ReferenceTaskName = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetRetryCount() float32 {
	if o == nil || o.RetryCount == nil {
		var ret float32
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetRetryCountOk() (*float32, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given float32 and assigns it to the RetryCount field.
func (o *WorkflowTaskNotificationAllOf) SetRetryCount(v float32) {
	o.RetryCount = &v
}

// GetScheduledTime returns the ScheduledTime field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetScheduledTime() string {
	if o == nil || o.ScheduledTime == nil {
		var ret string
		return ret
	}
	return *o.ScheduledTime
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetScheduledTimeOk() (*string, bool) {
	if o == nil || o.ScheduledTime == nil {
		return nil, false
	}
	return o.ScheduledTime, true
}

// HasScheduledTime returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasScheduledTime() bool {
	if o != nil && o.ScheduledTime != nil {
		return true
	}

	return false
}

// SetScheduledTime gets a reference to the given string and assigns it to the ScheduledTime field.
func (o *WorkflowTaskNotificationAllOf) SetScheduledTime(v string) {
	o.ScheduledTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *WorkflowTaskNotificationAllOf) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowTaskNotificationAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskDefName returns the TaskDefName field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetTaskDefName() string {
	if o == nil || o.TaskDefName == nil {
		var ret string
		return ret
	}
	return *o.TaskDefName
}

// GetTaskDefNameOk returns a tuple with the TaskDefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetTaskDefNameOk() (*string, bool) {
	if o == nil || o.TaskDefName == nil {
		return nil, false
	}
	return o.TaskDefName, true
}

// HasTaskDefName returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasTaskDefName() bool {
	if o != nil && o.TaskDefName != nil {
		return true
	}

	return false
}

// SetTaskDefName gets a reference to the given string and assigns it to the TaskDefName field.
func (o *WorkflowTaskNotificationAllOf) SetTaskDefName(v string) {
	o.TaskDefName = &v
}

// GetTaskDescription returns the TaskDescription field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetTaskDescription() string {
	if o == nil || o.TaskDescription == nil {
		var ret string
		return ret
	}
	return *o.TaskDescription
}

// GetTaskDescriptionOk returns a tuple with the TaskDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetTaskDescriptionOk() (*string, bool) {
	if o == nil || o.TaskDescription == nil {
		return nil, false
	}
	return o.TaskDescription, true
}

// HasTaskDescription returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasTaskDescription() bool {
	if o != nil && o.TaskDescription != nil {
		return true
	}

	return false
}

// SetTaskDescription gets a reference to the given string and assigns it to the TaskDescription field.
func (o *WorkflowTaskNotificationAllOf) SetTaskDescription(v string) {
	o.TaskDescription = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *WorkflowTaskNotificationAllOf) SetTaskId(v string) {
	o.TaskId = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetTaskType() string {
	if o == nil || o.TaskType == nil {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetTaskTypeOk() (*string, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *WorkflowTaskNotificationAllOf) SetTaskType(v string) {
	o.TaskType = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetUpdateTime() string {
	if o == nil || o.UpdateTime == nil {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetUpdateTimeOk() (*string, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *WorkflowTaskNotificationAllOf) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *WorkflowTaskNotificationAllOf) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetWorkflowTaskType returns the WorkflowTaskType field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowTaskType() string {
	if o == nil || o.WorkflowTaskType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowTaskType
}

// GetWorkflowTaskTypeOk returns a tuple with the WorkflowTaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowTaskTypeOk() (*string, bool) {
	if o == nil || o.WorkflowTaskType == nil {
		return nil, false
	}
	return o.WorkflowTaskType, true
}

// HasWorkflowTaskType returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasWorkflowTaskType() bool {
	if o != nil && o.WorkflowTaskType != nil {
		return true
	}

	return false
}

// SetWorkflowTaskType gets a reference to the given string and assigns it to the WorkflowTaskType field.
func (o *WorkflowTaskNotificationAllOf) SetWorkflowTaskType(v string) {
	o.WorkflowTaskType = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskNotificationAllOf) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowTaskNotificationAllOf) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowTaskNotificationAllOf) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o WorkflowTaskNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CorrelationId != nil {
		toSerialize["CorrelationId"] = o.CorrelationId
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.ReasonForIncompletion != nil {
		toSerialize["ReasonForIncompletion"] = o.ReasonForIncompletion
	}
	if o.ReferenceTaskName != nil {
		toSerialize["ReferenceTaskName"] = o.ReferenceTaskName
	}
	if o.RetryCount != nil {
		toSerialize["RetryCount"] = o.RetryCount
	}
	if o.ScheduledTime != nil {
		toSerialize["ScheduledTime"] = o.ScheduledTime
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskDefName != nil {
		toSerialize["TaskDefName"] = o.TaskDefName
	}
	if o.TaskDescription != nil {
		toSerialize["TaskDescription"] = o.TaskDescription
	}
	if o.TaskId != nil {
		toSerialize["TaskId"] = o.TaskId
	}
	if o.TaskType != nil {
		toSerialize["TaskType"] = o.TaskType
	}
	if o.UpdateTime != nil {
		toSerialize["UpdateTime"] = o.UpdateTime
	}
	if o.WorkflowId != nil {
		toSerialize["WorkflowId"] = o.WorkflowId
	}
	if o.WorkflowTaskType != nil {
		toSerialize["WorkflowTaskType"] = o.WorkflowTaskType
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskNotificationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTaskNotificationAllOf := _WorkflowTaskNotificationAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTaskNotificationAllOf); err == nil {
		*o = WorkflowTaskNotificationAllOf(varWorkflowTaskNotificationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CorrelationId")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "ReasonForIncompletion")
		delete(additionalProperties, "ReferenceTaskName")
		delete(additionalProperties, "RetryCount")
		delete(additionalProperties, "ScheduledTime")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskDefName")
		delete(additionalProperties, "TaskDescription")
		delete(additionalProperties, "TaskId")
		delete(additionalProperties, "TaskType")
		delete(additionalProperties, "UpdateTime")
		delete(additionalProperties, "WorkflowId")
		delete(additionalProperties, "WorkflowTaskType")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskNotificationAllOf struct {
	value *WorkflowTaskNotificationAllOf
	isSet bool
}

func (v NullableWorkflowTaskNotificationAllOf) Get() *WorkflowTaskNotificationAllOf {
	return v.value
}

func (v *NullableWorkflowTaskNotificationAllOf) Set(val *WorkflowTaskNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskNotificationAllOf(val *WorkflowTaskNotificationAllOf) *NullableWorkflowTaskNotificationAllOf {
	return &NullableWorkflowTaskNotificationAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTaskNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
