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

// OpenapiTaskGenerationResult Provides information about the status of the tasks created for a TaskGenerationRequest.
type OpenapiTaskGenerationResult struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An error message for when task creation fails.
	FailureReason *string `json:"FailureReason,omitempty"`
	// The display label of the task that is created.
	TaskDisplayName *string `json:"TaskDisplayName,omitempty"`
	// The name of the task that is created.
	TaskName *string `json:"TaskName,omitempty"`
	// Denotes the status of the task creation. * `none` - Indicates the default status * `InProgress` - Indicates that operation is in progress * `Completed` - Indicates that the operation is complete * `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details.
	TaskStatus *string `json:"TaskStatus,omitempty"`
	// The version number of the created tasks.
	TaskVersion           *int64                                    `json:"TaskVersion,omitempty"`
	TaskGenerationRequest *OpenapiTaskGenerationRequestRelationship `json:"TaskGenerationRequest,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _OpenapiTaskGenerationResult OpenapiTaskGenerationResult

// NewOpenapiTaskGenerationResult instantiates a new OpenapiTaskGenerationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiTaskGenerationResult(classId string, objectType string) *OpenapiTaskGenerationResult {
	this := OpenapiTaskGenerationResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOpenapiTaskGenerationResultWithDefaults instantiates a new OpenapiTaskGenerationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiTaskGenerationResultWithDefaults() *OpenapiTaskGenerationResult {
	this := OpenapiTaskGenerationResult{}
	var classId string = "openapi.TaskGenerationResult"
	this.ClassId = classId
	var objectType string = "openapi.TaskGenerationResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiTaskGenerationResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiTaskGenerationResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiTaskGenerationResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiTaskGenerationResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *OpenapiTaskGenerationResult) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetTaskDisplayName returns the TaskDisplayName field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetTaskDisplayName() string {
	if o == nil || o.TaskDisplayName == nil {
		var ret string
		return ret
	}
	return *o.TaskDisplayName
}

// GetTaskDisplayNameOk returns a tuple with the TaskDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetTaskDisplayNameOk() (*string, bool) {
	if o == nil || o.TaskDisplayName == nil {
		return nil, false
	}
	return o.TaskDisplayName, true
}

// HasTaskDisplayName returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasTaskDisplayName() bool {
	if o != nil && o.TaskDisplayName != nil {
		return true
	}

	return false
}

// SetTaskDisplayName gets a reference to the given string and assigns it to the TaskDisplayName field.
func (o *OpenapiTaskGenerationResult) SetTaskDisplayName(v string) {
	o.TaskDisplayName = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetTaskName() string {
	if o == nil || o.TaskName == nil {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetTaskNameOk() (*string, bool) {
	if o == nil || o.TaskName == nil {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasTaskName() bool {
	if o != nil && o.TaskName != nil {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *OpenapiTaskGenerationResult) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetTaskStatus() string {
	if o == nil || o.TaskStatus == nil {
		var ret string
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetTaskStatusOk() (*string, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given string and assigns it to the TaskStatus field.
func (o *OpenapiTaskGenerationResult) SetTaskStatus(v string) {
	o.TaskStatus = &v
}

// GetTaskVersion returns the TaskVersion field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetTaskVersion() int64 {
	if o == nil || o.TaskVersion == nil {
		var ret int64
		return ret
	}
	return *o.TaskVersion
}

// GetTaskVersionOk returns a tuple with the TaskVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetTaskVersionOk() (*int64, bool) {
	if o == nil || o.TaskVersion == nil {
		return nil, false
	}
	return o.TaskVersion, true
}

// HasTaskVersion returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasTaskVersion() bool {
	if o != nil && o.TaskVersion != nil {
		return true
	}

	return false
}

// SetTaskVersion gets a reference to the given int64 and assigns it to the TaskVersion field.
func (o *OpenapiTaskGenerationResult) SetTaskVersion(v int64) {
	o.TaskVersion = &v
}

// GetTaskGenerationRequest returns the TaskGenerationRequest field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationResult) GetTaskGenerationRequest() OpenapiTaskGenerationRequestRelationship {
	if o == nil || o.TaskGenerationRequest == nil {
		var ret OpenapiTaskGenerationRequestRelationship
		return ret
	}
	return *o.TaskGenerationRequest
}

// GetTaskGenerationRequestOk returns a tuple with the TaskGenerationRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationResult) GetTaskGenerationRequestOk() (*OpenapiTaskGenerationRequestRelationship, bool) {
	if o == nil || o.TaskGenerationRequest == nil {
		return nil, false
	}
	return o.TaskGenerationRequest, true
}

// HasTaskGenerationRequest returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationResult) HasTaskGenerationRequest() bool {
	if o != nil && o.TaskGenerationRequest != nil {
		return true
	}

	return false
}

// SetTaskGenerationRequest gets a reference to the given OpenapiTaskGenerationRequestRelationship and assigns it to the TaskGenerationRequest field.
func (o *OpenapiTaskGenerationResult) SetTaskGenerationRequest(v OpenapiTaskGenerationRequestRelationship) {
	o.TaskGenerationRequest = &v
}

func (o OpenapiTaskGenerationResult) MarshalJSON() ([]byte, error) {
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
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.TaskDisplayName != nil {
		toSerialize["TaskDisplayName"] = o.TaskDisplayName
	}
	if o.TaskName != nil {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.TaskStatus != nil {
		toSerialize["TaskStatus"] = o.TaskStatus
	}
	if o.TaskVersion != nil {
		toSerialize["TaskVersion"] = o.TaskVersion
	}
	if o.TaskGenerationRequest != nil {
		toSerialize["TaskGenerationRequest"] = o.TaskGenerationRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiTaskGenerationResult) UnmarshalJSON(bytes []byte) (err error) {
	type OpenapiTaskGenerationResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An error message for when task creation fails.
		FailureReason *string `json:"FailureReason,omitempty"`
		// The display label of the task that is created.
		TaskDisplayName *string `json:"TaskDisplayName,omitempty"`
		// The name of the task that is created.
		TaskName *string `json:"TaskName,omitempty"`
		// Denotes the status of the task creation. * `none` - Indicates the default status * `InProgress` - Indicates that operation is in progress * `Completed` - Indicates that the operation is complete * `Failed` - Indicates that the operation has failed. Check the failureReason attribute for more details.
		TaskStatus *string `json:"TaskStatus,omitempty"`
		// The version number of the created tasks.
		TaskVersion           *int64                                    `json:"TaskVersion,omitempty"`
		TaskGenerationRequest *OpenapiTaskGenerationRequestRelationship `json:"TaskGenerationRequest,omitempty"`
	}

	varOpenapiTaskGenerationResultWithoutEmbeddedStruct := OpenapiTaskGenerationResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOpenapiTaskGenerationResultWithoutEmbeddedStruct)
	if err == nil {
		varOpenapiTaskGenerationResult := _OpenapiTaskGenerationResult{}
		varOpenapiTaskGenerationResult.ClassId = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.ClassId
		varOpenapiTaskGenerationResult.ObjectType = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.ObjectType
		varOpenapiTaskGenerationResult.FailureReason = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.FailureReason
		varOpenapiTaskGenerationResult.TaskDisplayName = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.TaskDisplayName
		varOpenapiTaskGenerationResult.TaskName = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.TaskName
		varOpenapiTaskGenerationResult.TaskStatus = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.TaskStatus
		varOpenapiTaskGenerationResult.TaskVersion = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.TaskVersion
		varOpenapiTaskGenerationResult.TaskGenerationRequest = varOpenapiTaskGenerationResultWithoutEmbeddedStruct.TaskGenerationRequest
		*o = OpenapiTaskGenerationResult(varOpenapiTaskGenerationResult)
	} else {
		return err
	}

	varOpenapiTaskGenerationResult := _OpenapiTaskGenerationResult{}

	err = json.Unmarshal(bytes, &varOpenapiTaskGenerationResult)
	if err == nil {
		o.MoBaseMo = varOpenapiTaskGenerationResult.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "TaskDisplayName")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "TaskStatus")
		delete(additionalProperties, "TaskVersion")
		delete(additionalProperties, "TaskGenerationRequest")

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

type NullableOpenapiTaskGenerationResult struct {
	value *OpenapiTaskGenerationResult
	isSet bool
}

func (v NullableOpenapiTaskGenerationResult) Get() *OpenapiTaskGenerationResult {
	return v.value
}

func (v *NullableOpenapiTaskGenerationResult) Set(val *OpenapiTaskGenerationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiTaskGenerationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiTaskGenerationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiTaskGenerationResult(val *OpenapiTaskGenerationResult) *NullableOpenapiTaskGenerationResult {
	return &NullableOpenapiTaskGenerationResult{value: val, isSet: true}
}

func (v NullableOpenapiTaskGenerationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiTaskGenerationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
