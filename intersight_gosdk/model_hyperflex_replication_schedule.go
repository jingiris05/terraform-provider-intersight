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

// checks if the HyperflexReplicationSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexReplicationSchedule{}

// HyperflexReplicationSchedule The HyperFlex Cluster backup policy replication schedule.
type HyperflexReplicationSchedule struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Time interval between two copies in minutes.
	BackupInterval       *int64 `json:"BackupInterval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationSchedule HyperflexReplicationSchedule

// NewHyperflexReplicationSchedule instantiates a new HyperflexReplicationSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationSchedule(classId string, objectType string) *HyperflexReplicationSchedule {
	this := HyperflexReplicationSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	var backupInterval int64 = 240
	this.BackupInterval = &backupInterval
	return &this
}

// NewHyperflexReplicationScheduleWithDefaults instantiates a new HyperflexReplicationSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationScheduleWithDefaults() *HyperflexReplicationSchedule {
	this := HyperflexReplicationSchedule{}
	var classId string = "hyperflex.ReplicationSchedule"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationSchedule"
	this.ObjectType = objectType
	var backupInterval int64 = 240
	this.BackupInterval = &backupInterval
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ReplicationSchedule" of the ClassId field.
func (o *HyperflexReplicationSchedule) GetDefaultClassId() interface{} {
	return "hyperflex.ReplicationSchedule"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ReplicationSchedule" of the ObjectType field.
func (o *HyperflexReplicationSchedule) GetDefaultObjectType() interface{} {
	return "hyperflex.ReplicationSchedule"
}

// GetBackupInterval returns the BackupInterval field value if set, zero value otherwise.
func (o *HyperflexReplicationSchedule) GetBackupInterval() int64 {
	if o == nil || IsNil(o.BackupInterval) {
		var ret int64
		return ret
	}
	return *o.BackupInterval
}

// GetBackupIntervalOk returns a tuple with the BackupInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationSchedule) GetBackupIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.BackupInterval) {
		return nil, false
	}
	return o.BackupInterval, true
}

// HasBackupInterval returns a boolean if a field has been set.
func (o *HyperflexReplicationSchedule) HasBackupInterval() bool {
	if o != nil && !IsNil(o.BackupInterval) {
		return true
	}

	return false
}

// SetBackupInterval gets a reference to the given int64 and assigns it to the BackupInterval field.
func (o *HyperflexReplicationSchedule) SetBackupInterval(v int64) {
	o.BackupInterval = &v
}

func (o HyperflexReplicationSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexReplicationSchedule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BackupInterval) {
		toSerialize["BackupInterval"] = o.BackupInterval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexReplicationSchedule) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexReplicationScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Time interval between two copies in minutes.
		BackupInterval *int64 `json:"BackupInterval,omitempty"`
	}

	varHyperflexReplicationScheduleWithoutEmbeddedStruct := HyperflexReplicationScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexReplicationScheduleWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexReplicationSchedule := _HyperflexReplicationSchedule{}
		varHyperflexReplicationSchedule.ClassId = varHyperflexReplicationScheduleWithoutEmbeddedStruct.ClassId
		varHyperflexReplicationSchedule.ObjectType = varHyperflexReplicationScheduleWithoutEmbeddedStruct.ObjectType
		varHyperflexReplicationSchedule.BackupInterval = varHyperflexReplicationScheduleWithoutEmbeddedStruct.BackupInterval
		*o = HyperflexReplicationSchedule(varHyperflexReplicationSchedule)
	} else {
		return err
	}

	varHyperflexReplicationSchedule := _HyperflexReplicationSchedule{}

	err = json.Unmarshal(data, &varHyperflexReplicationSchedule)
	if err == nil {
		o.MoBaseComplexType = varHyperflexReplicationSchedule.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupInterval")

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

type NullableHyperflexReplicationSchedule struct {
	value *HyperflexReplicationSchedule
	isSet bool
}

func (v NullableHyperflexReplicationSchedule) Get() *HyperflexReplicationSchedule {
	return v.value
}

func (v *NullableHyperflexReplicationSchedule) Set(val *HyperflexReplicationSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationSchedule(val *HyperflexReplicationSchedule) *NullableHyperflexReplicationSchedule {
	return &NullableHyperflexReplicationSchedule{value: val, isSet: true}
}

func (v NullableHyperflexReplicationSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
