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
	"reflect"
	"strings"
)

// HyperflexReplicationStatus Replication status information for this particular snapshot.
type HyperflexReplicationStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of bytes currently replicated.
	BytesReplicated *int64 `json:"BytesReplicated,omitempty"`
	// Replication end time for this snapshot.
	EndTime             *int64                           `json:"EndTime,omitempty"`
	Error               NullableHyperflexErrorStack      `json:"Error,omitempty"`
	PackEntityReference NullableHyperflexEntityReference `json:"PackEntityReference,omitempty"`
	// Completion percentage for the replication job.
	PctComplete *int64                     `json:"PctComplete,omitempty"`
	RpoStatus   NullableHyperflexRpoStatus `json:"RpoStatus,omitempty"`
	// Replication start time for this snapshot.
	StartTime *int64 `json:"StartTime,omitempty"`
	// Current replication state for a particular snapshot. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationStatus HyperflexReplicationStatus

// NewHyperflexReplicationStatus instantiates a new HyperflexReplicationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationStatus(classId string, objectType string) *HyperflexReplicationStatus {
	this := HyperflexReplicationStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "NONE"
	this.Status = &status
	return &this
}

// NewHyperflexReplicationStatusWithDefaults instantiates a new HyperflexReplicationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationStatusWithDefaults() *HyperflexReplicationStatus {
	this := HyperflexReplicationStatus{}
	var classId string = "hyperflex.ReplicationStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationStatus"
	this.ObjectType = objectType
	var status string = "NONE"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBytesReplicated returns the BytesReplicated field value if set, zero value otherwise.
func (o *HyperflexReplicationStatus) GetBytesReplicated() int64 {
	if o == nil || o.BytesReplicated == nil {
		var ret int64
		return ret
	}
	return *o.BytesReplicated
}

// GetBytesReplicatedOk returns a tuple with the BytesReplicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetBytesReplicatedOk() (*int64, bool) {
	if o == nil || o.BytesReplicated == nil {
		return nil, false
	}
	return o.BytesReplicated, true
}

// HasBytesReplicated returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasBytesReplicated() bool {
	if o != nil && o.BytesReplicated != nil {
		return true
	}

	return false
}

// SetBytesReplicated gets a reference to the given int64 and assigns it to the BytesReplicated field.
func (o *HyperflexReplicationStatus) SetBytesReplicated(v int64) {
	o.BytesReplicated = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *HyperflexReplicationStatus) GetEndTime() int64 {
	if o == nil || o.EndTime == nil {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetEndTimeOk() (*int64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *HyperflexReplicationStatus) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationStatus) GetError() HyperflexErrorStack {
	if o == nil || o.Error.Get() == nil {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationStatus) GetErrorOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableHyperflexErrorStack and assigns it to the Error field.
func (o *HyperflexReplicationStatus) SetError(v HyperflexErrorStack) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *HyperflexReplicationStatus) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HyperflexReplicationStatus) UnsetError() {
	o.Error.Unset()
}

// GetPackEntityReference returns the PackEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationStatus) GetPackEntityReference() HyperflexEntityReference {
	if o == nil || o.PackEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.PackEntityReference.Get()
}

// GetPackEntityReferenceOk returns a tuple with the PackEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationStatus) GetPackEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackEntityReference.Get(), o.PackEntityReference.IsSet()
}

// HasPackEntityReference returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasPackEntityReference() bool {
	if o != nil && o.PackEntityReference.IsSet() {
		return true
	}

	return false
}

// SetPackEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the PackEntityReference field.
func (o *HyperflexReplicationStatus) SetPackEntityReference(v HyperflexEntityReference) {
	o.PackEntityReference.Set(&v)
}

// SetPackEntityReferenceNil sets the value for PackEntityReference to be an explicit nil
func (o *HyperflexReplicationStatus) SetPackEntityReferenceNil() {
	o.PackEntityReference.Set(nil)
}

// UnsetPackEntityReference ensures that no value is present for PackEntityReference, not even an explicit nil
func (o *HyperflexReplicationStatus) UnsetPackEntityReference() {
	o.PackEntityReference.Unset()
}

// GetPctComplete returns the PctComplete field value if set, zero value otherwise.
func (o *HyperflexReplicationStatus) GetPctComplete() int64 {
	if o == nil || o.PctComplete == nil {
		var ret int64
		return ret
	}
	return *o.PctComplete
}

// GetPctCompleteOk returns a tuple with the PctComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetPctCompleteOk() (*int64, bool) {
	if o == nil || o.PctComplete == nil {
		return nil, false
	}
	return o.PctComplete, true
}

// HasPctComplete returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasPctComplete() bool {
	if o != nil && o.PctComplete != nil {
		return true
	}

	return false
}

// SetPctComplete gets a reference to the given int64 and assigns it to the PctComplete field.
func (o *HyperflexReplicationStatus) SetPctComplete(v int64) {
	o.PctComplete = &v
}

// GetRpoStatus returns the RpoStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationStatus) GetRpoStatus() HyperflexRpoStatus {
	if o == nil || o.RpoStatus.Get() == nil {
		var ret HyperflexRpoStatus
		return ret
	}
	return *o.RpoStatus.Get()
}

// GetRpoStatusOk returns a tuple with the RpoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationStatus) GetRpoStatusOk() (*HyperflexRpoStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpoStatus.Get(), o.RpoStatus.IsSet()
}

// HasRpoStatus returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasRpoStatus() bool {
	if o != nil && o.RpoStatus.IsSet() {
		return true
	}

	return false
}

// SetRpoStatus gets a reference to the given NullableHyperflexRpoStatus and assigns it to the RpoStatus field.
func (o *HyperflexReplicationStatus) SetRpoStatus(v HyperflexRpoStatus) {
	o.RpoStatus.Set(&v)
}

// SetRpoStatusNil sets the value for RpoStatus to be an explicit nil
func (o *HyperflexReplicationStatus) SetRpoStatusNil() {
	o.RpoStatus.Set(nil)
}

// UnsetRpoStatus ensures that no value is present for RpoStatus, not even an explicit nil
func (o *HyperflexReplicationStatus) UnsetRpoStatus() {
	o.RpoStatus.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *HyperflexReplicationStatus) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *HyperflexReplicationStatus) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexReplicationStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexReplicationStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexReplicationStatus) SetStatus(v string) {
	o.Status = &v
}

func (o HyperflexReplicationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BytesReplicated != nil {
		toSerialize["BytesReplicated"] = o.BytesReplicated
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Error.IsSet() {
		toSerialize["Error"] = o.Error.Get()
	}
	if o.PackEntityReference.IsSet() {
		toSerialize["PackEntityReference"] = o.PackEntityReference.Get()
	}
	if o.PctComplete != nil {
		toSerialize["PctComplete"] = o.PctComplete
	}
	if o.RpoStatus.IsSet() {
		toSerialize["RpoStatus"] = o.RpoStatus.Get()
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexReplicationStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of bytes currently replicated.
		BytesReplicated *int64 `json:"BytesReplicated,omitempty"`
		// Replication end time for this snapshot.
		EndTime             *int64                           `json:"EndTime,omitempty"`
		Error               NullableHyperflexErrorStack      `json:"Error,omitempty"`
		PackEntityReference NullableHyperflexEntityReference `json:"PackEntityReference,omitempty"`
		// Completion percentage for the replication job.
		PctComplete *int64                     `json:"PctComplete,omitempty"`
		RpoStatus   NullableHyperflexRpoStatus `json:"RpoStatus,omitempty"`
		// Replication start time for this snapshot.
		StartTime *int64 `json:"StartTime,omitempty"`
		// Current replication state for a particular snapshot. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
		Status *string `json:"Status,omitempty"`
	}

	varHyperflexReplicationStatusWithoutEmbeddedStruct := HyperflexReplicationStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexReplicationStatusWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexReplicationStatus := _HyperflexReplicationStatus{}
		varHyperflexReplicationStatus.ClassId = varHyperflexReplicationStatusWithoutEmbeddedStruct.ClassId
		varHyperflexReplicationStatus.ObjectType = varHyperflexReplicationStatusWithoutEmbeddedStruct.ObjectType
		varHyperflexReplicationStatus.BytesReplicated = varHyperflexReplicationStatusWithoutEmbeddedStruct.BytesReplicated
		varHyperflexReplicationStatus.EndTime = varHyperflexReplicationStatusWithoutEmbeddedStruct.EndTime
		varHyperflexReplicationStatus.Error = varHyperflexReplicationStatusWithoutEmbeddedStruct.Error
		varHyperflexReplicationStatus.PackEntityReference = varHyperflexReplicationStatusWithoutEmbeddedStruct.PackEntityReference
		varHyperflexReplicationStatus.PctComplete = varHyperflexReplicationStatusWithoutEmbeddedStruct.PctComplete
		varHyperflexReplicationStatus.RpoStatus = varHyperflexReplicationStatusWithoutEmbeddedStruct.RpoStatus
		varHyperflexReplicationStatus.StartTime = varHyperflexReplicationStatusWithoutEmbeddedStruct.StartTime
		varHyperflexReplicationStatus.Status = varHyperflexReplicationStatusWithoutEmbeddedStruct.Status
		*o = HyperflexReplicationStatus(varHyperflexReplicationStatus)
	} else {
		return err
	}

	varHyperflexReplicationStatus := _HyperflexReplicationStatus{}

	err = json.Unmarshal(bytes, &varHyperflexReplicationStatus)
	if err == nil {
		o.MoBaseComplexType = varHyperflexReplicationStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BytesReplicated")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "PackEntityReference")
		delete(additionalProperties, "PctComplete")
		delete(additionalProperties, "RpoStatus")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")

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

type NullableHyperflexReplicationStatus struct {
	value *HyperflexReplicationStatus
	isSet bool
}

func (v NullableHyperflexReplicationStatus) Get() *HyperflexReplicationStatus {
	return v.value
}

func (v *NullableHyperflexReplicationStatus) Set(val *HyperflexReplicationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationStatus(val *HyperflexReplicationStatus) *NullableHyperflexReplicationStatus {
	return &NullableHyperflexReplicationStatus{value: val, isSet: true}
}

func (v NullableHyperflexReplicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
