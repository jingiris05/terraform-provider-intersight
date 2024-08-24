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

// checks if the StorageBaseReplicationSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageBaseReplicationSchedule{}

// StorageBaseReplicationSchedule Configuration parameters for snapshot creation on target arrays.
type StorageBaseReplicationSchedule struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Replication frequency. It is an interval at which replication is set to trigger. Examples:     PT2H, Snapshot is generated every 2 hours.     P30D, Snapshot is scheduled for every 30 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	Frequency *string `json:"Frequency,omitempty"`
	// Replication schedule name.
	Name *string `json:"Name,omitempty"`
	// Preferred time of day at which to replicate the snapshots on target array. It is applicable only if the replication frequency is set for a day or more. Format: hh:mm:ss Example: 15:00:00, Replication is set for 3:00 PM.
	ReplicationTime *string `json:"ReplicationTime,omitempty"`
	// Duration to keep the replicated snapshots on the targets. Replicated snapshots are deleted from target array once the retention period has elapsed. Examples: P30D, Snapshots are available for 30 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	RetentionTime        *string `json:"RetentionTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseReplicationSchedule StorageBaseReplicationSchedule

// NewStorageBaseReplicationSchedule instantiates a new StorageBaseReplicationSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseReplicationSchedule(classId string, objectType string) *StorageBaseReplicationSchedule {
	this := StorageBaseReplicationSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseReplicationScheduleWithDefaults instantiates a new StorageBaseReplicationSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseReplicationScheduleWithDefaults() *StorageBaseReplicationSchedule {
	this := StorageBaseReplicationSchedule{}
	var classId string = "storage.PureReplicationSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseReplicationSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseReplicationSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureReplicationSchedule" of the ClassId field.
func (o *StorageBaseReplicationSchedule) GetDefaultClassId() interface{} {
	return "storage.PureReplicationSchedule"
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseReplicationSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseReplicationSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureReplicationSchedule" of the ObjectType field.
func (o *StorageBaseReplicationSchedule) GetDefaultObjectType() interface{} {
	return "storage.PureReplicationSchedule"
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *StorageBaseReplicationSchedule) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *StorageBaseReplicationSchedule) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *StorageBaseReplicationSchedule) SetFrequency(v string) {
	o.Frequency = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseReplicationSchedule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseReplicationSchedule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseReplicationSchedule) SetName(v string) {
	o.Name = &v
}

// GetReplicationTime returns the ReplicationTime field value if set, zero value otherwise.
func (o *StorageBaseReplicationSchedule) GetReplicationTime() string {
	if o == nil || IsNil(o.ReplicationTime) {
		var ret string
		return ret
	}
	return *o.ReplicationTime
}

// GetReplicationTimeOk returns a tuple with the ReplicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetReplicationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationTime) {
		return nil, false
	}
	return o.ReplicationTime, true
}

// HasReplicationTime returns a boolean if a field has been set.
func (o *StorageBaseReplicationSchedule) HasReplicationTime() bool {
	if o != nil && !IsNil(o.ReplicationTime) {
		return true
	}

	return false
}

// SetReplicationTime gets a reference to the given string and assigns it to the ReplicationTime field.
func (o *StorageBaseReplicationSchedule) SetReplicationTime(v string) {
	o.ReplicationTime = &v
}

// GetRetentionTime returns the RetentionTime field value if set, zero value otherwise.
func (o *StorageBaseReplicationSchedule) GetRetentionTime() string {
	if o == nil || IsNil(o.RetentionTime) {
		var ret string
		return ret
	}
	return *o.RetentionTime
}

// GetRetentionTimeOk returns a tuple with the RetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationSchedule) GetRetentionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionTime) {
		return nil, false
	}
	return o.RetentionTime, true
}

// HasRetentionTime returns a boolean if a field has been set.
func (o *StorageBaseReplicationSchedule) HasRetentionTime() bool {
	if o != nil && !IsNil(o.RetentionTime) {
		return true
	}

	return false
}

// SetRetentionTime gets a reference to the given string and assigns it to the RetentionTime field.
func (o *StorageBaseReplicationSchedule) SetRetentionTime(v string) {
	o.RetentionTime = &v
}

func (o StorageBaseReplicationSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageBaseReplicationSchedule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Frequency) {
		toSerialize["Frequency"] = o.Frequency
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ReplicationTime) {
		toSerialize["ReplicationTime"] = o.ReplicationTime
	}
	if !IsNil(o.RetentionTime) {
		toSerialize["RetentionTime"] = o.RetentionTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageBaseReplicationSchedule) UnmarshalJSON(data []byte) (err error) {
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
	type StorageBaseReplicationScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Replication frequency. It is an interval at which replication is set to trigger. Examples:     PT2H, Snapshot is generated every 2 hours.     P30D, Snapshot is scheduled for every 30 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
		Frequency *string `json:"Frequency,omitempty"`
		// Replication schedule name.
		Name *string `json:"Name,omitempty"`
		// Preferred time of day at which to replicate the snapshots on target array. It is applicable only if the replication frequency is set for a day or more. Format: hh:mm:ss Example: 15:00:00, Replication is set for 3:00 PM.
		ReplicationTime *string `json:"ReplicationTime,omitempty"`
		// Duration to keep the replicated snapshots on the targets. Replicated snapshots are deleted from target array once the retention period has elapsed. Examples: P30D, Snapshots are available for 30 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
		RetentionTime *string `json:"RetentionTime,omitempty"`
	}

	varStorageBaseReplicationScheduleWithoutEmbeddedStruct := StorageBaseReplicationScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageBaseReplicationScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseReplicationSchedule := _StorageBaseReplicationSchedule{}
		varStorageBaseReplicationSchedule.ClassId = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.ClassId
		varStorageBaseReplicationSchedule.ObjectType = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.ObjectType
		varStorageBaseReplicationSchedule.Frequency = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.Frequency
		varStorageBaseReplicationSchedule.Name = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.Name
		varStorageBaseReplicationSchedule.ReplicationTime = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.ReplicationTime
		varStorageBaseReplicationSchedule.RetentionTime = varStorageBaseReplicationScheduleWithoutEmbeddedStruct.RetentionTime
		*o = StorageBaseReplicationSchedule(varStorageBaseReplicationSchedule)
	} else {
		return err
	}

	varStorageBaseReplicationSchedule := _StorageBaseReplicationSchedule{}

	err = json.Unmarshal(data, &varStorageBaseReplicationSchedule)
	if err == nil {
		o.MoBaseMo = varStorageBaseReplicationSchedule.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Frequency")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ReplicationTime")
		delete(additionalProperties, "RetentionTime")

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

type NullableStorageBaseReplicationSchedule struct {
	value *StorageBaseReplicationSchedule
	isSet bool
}

func (v NullableStorageBaseReplicationSchedule) Get() *StorageBaseReplicationSchedule {
	return v.value
}

func (v *NullableStorageBaseReplicationSchedule) Set(val *StorageBaseReplicationSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseReplicationSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseReplicationSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseReplicationSchedule(val *StorageBaseReplicationSchedule) *NullableStorageBaseReplicationSchedule {
	return &NullableStorageBaseReplicationSchedule{value: val, isSet: true}
}

func (v NullableStorageBaseReplicationSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseReplicationSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
