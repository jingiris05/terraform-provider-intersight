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

// checks if the StoragePureVolumeSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePureVolumeSnapshot{}

// StoragePureVolumeSnapshot Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
type StoragePureVolumeSnapshot struct {
	StorageBaseSnapshot
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique serial number of the snapshot allocated by the storage array.
	Serial                  *string                                                `json:"Serial,omitempty"`
	Array                   NullableStoragePureArrayRelationship                   `json:"Array,omitempty"`
	ProtectionGroupSnapshot NullableStoragePureProtectionGroupSnapshotRelationship `json:"ProtectionGroupSnapshot,omitempty"`
	RegisteredDevice        NullableAssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	Volume                  NullableStoragePureVolumeRelationship                  `json:"Volume,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _StoragePureVolumeSnapshot StoragePureVolumeSnapshot

// NewStoragePureVolumeSnapshot instantiates a new StoragePureVolumeSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureVolumeSnapshot(classId string, objectType string) *StoragePureVolumeSnapshot {
	this := StoragePureVolumeSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureVolumeSnapshotWithDefaults instantiates a new StoragePureVolumeSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureVolumeSnapshotWithDefaults() *StoragePureVolumeSnapshot {
	this := StoragePureVolumeSnapshot{}
	var classId string = "storage.PureVolumeSnapshot"
	this.ClassId = classId
	var objectType string = "storage.PureVolumeSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureVolumeSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureVolumeSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PureVolumeSnapshot" of the ClassId field.
func (o *StoragePureVolumeSnapshot) GetDefaultClassId() interface{} {
	return "storage.PureVolumeSnapshot"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureVolumeSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureVolumeSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PureVolumeSnapshot" of the ObjectType field.
func (o *StoragePureVolumeSnapshot) GetDefaultObjectType() interface{} {
	return "storage.PureVolumeSnapshot"
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *StoragePureVolumeSnapshot) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureVolumeSnapshot) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *StoragePureVolumeSnapshot) SetSerial(v string) {
	o.Serial = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureVolumeSnapshot) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureVolumeSnapshot) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureVolumeSnapshot) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePureVolumeSnapshot) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePureVolumeSnapshot) UnsetArray() {
	o.Array.Unset()
}

// GetProtectionGroupSnapshot returns the ProtectionGroupSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureVolumeSnapshot) GetProtectionGroupSnapshot() StoragePureProtectionGroupSnapshotRelationship {
	if o == nil || IsNil(o.ProtectionGroupSnapshot.Get()) {
		var ret StoragePureProtectionGroupSnapshotRelationship
		return ret
	}
	return *o.ProtectionGroupSnapshot.Get()
}

// GetProtectionGroupSnapshotOk returns a tuple with the ProtectionGroupSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureVolumeSnapshot) GetProtectionGroupSnapshotOk() (*StoragePureProtectionGroupSnapshotRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProtectionGroupSnapshot.Get(), o.ProtectionGroupSnapshot.IsSet()
}

// HasProtectionGroupSnapshot returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasProtectionGroupSnapshot() bool {
	if o != nil && o.ProtectionGroupSnapshot.IsSet() {
		return true
	}

	return false
}

// SetProtectionGroupSnapshot gets a reference to the given NullableStoragePureProtectionGroupSnapshotRelationship and assigns it to the ProtectionGroupSnapshot field.
func (o *StoragePureVolumeSnapshot) SetProtectionGroupSnapshot(v StoragePureProtectionGroupSnapshotRelationship) {
	o.ProtectionGroupSnapshot.Set(&v)
}

// SetProtectionGroupSnapshotNil sets the value for ProtectionGroupSnapshot to be an explicit nil
func (o *StoragePureVolumeSnapshot) SetProtectionGroupSnapshotNil() {
	o.ProtectionGroupSnapshot.Set(nil)
}

// UnsetProtectionGroupSnapshot ensures that no value is present for ProtectionGroupSnapshot, not even an explicit nil
func (o *StoragePureVolumeSnapshot) UnsetProtectionGroupSnapshot() {
	o.ProtectionGroupSnapshot.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureVolumeSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureVolumeSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureVolumeSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePureVolumeSnapshot) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePureVolumeSnapshot) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetVolume returns the Volume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureVolumeSnapshot) GetVolume() StoragePureVolumeRelationship {
	if o == nil || IsNil(o.Volume.Get()) {
		var ret StoragePureVolumeRelationship
		return ret
	}
	return *o.Volume.Get()
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureVolumeSnapshot) GetVolumeOk() (*StoragePureVolumeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volume.Get(), o.Volume.IsSet()
}

// HasVolume returns a boolean if a field has been set.
func (o *StoragePureVolumeSnapshot) HasVolume() bool {
	if o != nil && o.Volume.IsSet() {
		return true
	}

	return false
}

// SetVolume gets a reference to the given NullableStoragePureVolumeRelationship and assigns it to the Volume field.
func (o *StoragePureVolumeSnapshot) SetVolume(v StoragePureVolumeRelationship) {
	o.Volume.Set(&v)
}

// SetVolumeNil sets the value for Volume to be an explicit nil
func (o *StoragePureVolumeSnapshot) SetVolumeNil() {
	o.Volume.Set(nil)
}

// UnsetVolume ensures that no value is present for Volume, not even an explicit nil
func (o *StoragePureVolumeSnapshot) UnsetVolume() {
	o.Volume.Unset()
}

func (o StoragePureVolumeSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePureVolumeSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseSnapshot, errStorageBaseSnapshot := json.Marshal(o.StorageBaseSnapshot)
	if errStorageBaseSnapshot != nil {
		return map[string]interface{}{}, errStorageBaseSnapshot
	}
	errStorageBaseSnapshot = json.Unmarshal([]byte(serializedStorageBaseSnapshot), &toSerialize)
	if errStorageBaseSnapshot != nil {
		return map[string]interface{}{}, errStorageBaseSnapshot
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Serial) {
		toSerialize["Serial"] = o.Serial
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.ProtectionGroupSnapshot.IsSet() {
		toSerialize["ProtectionGroupSnapshot"] = o.ProtectionGroupSnapshot.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Volume.IsSet() {
		toSerialize["Volume"] = o.Volume.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePureVolumeSnapshot) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePureVolumeSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique serial number of the snapshot allocated by the storage array.
		Serial                  *string                                                `json:"Serial,omitempty"`
		Array                   NullableStoragePureArrayRelationship                   `json:"Array,omitempty"`
		ProtectionGroupSnapshot NullableStoragePureProtectionGroupSnapshotRelationship `json:"ProtectionGroupSnapshot,omitempty"`
		RegisteredDevice        NullableAssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
		Volume                  NullableStoragePureVolumeRelationship                  `json:"Volume,omitempty"`
	}

	varStoragePureVolumeSnapshotWithoutEmbeddedStruct := StoragePureVolumeSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePureVolumeSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureVolumeSnapshot := _StoragePureVolumeSnapshot{}
		varStoragePureVolumeSnapshot.ClassId = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.ClassId
		varStoragePureVolumeSnapshot.ObjectType = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.ObjectType
		varStoragePureVolumeSnapshot.Serial = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.Serial
		varStoragePureVolumeSnapshot.Array = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.Array
		varStoragePureVolumeSnapshot.ProtectionGroupSnapshot = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.ProtectionGroupSnapshot
		varStoragePureVolumeSnapshot.RegisteredDevice = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.RegisteredDevice
		varStoragePureVolumeSnapshot.Volume = varStoragePureVolumeSnapshotWithoutEmbeddedStruct.Volume
		*o = StoragePureVolumeSnapshot(varStoragePureVolumeSnapshot)
	} else {
		return err
	}

	varStoragePureVolumeSnapshot := _StoragePureVolumeSnapshot{}

	err = json.Unmarshal(data, &varStoragePureVolumeSnapshot)
	if err == nil {
		o.StorageBaseSnapshot = varStoragePureVolumeSnapshot.StorageBaseSnapshot
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroupSnapshot")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volume")

		// remove fields from embedded structs
		reflectStorageBaseSnapshot := reflect.ValueOf(o.StorageBaseSnapshot)
		for i := 0; i < reflectStorageBaseSnapshot.Type().NumField(); i++ {
			t := reflectStorageBaseSnapshot.Type().Field(i)

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

type NullableStoragePureVolumeSnapshot struct {
	value *StoragePureVolumeSnapshot
	isSet bool
}

func (v NullableStoragePureVolumeSnapshot) Get() *StoragePureVolumeSnapshot {
	return v.value
}

func (v *NullableStoragePureVolumeSnapshot) Set(val *StoragePureVolumeSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureVolumeSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureVolumeSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureVolumeSnapshot(val *StoragePureVolumeSnapshot) *NullableStoragePureVolumeSnapshot {
	return &NullableStoragePureVolumeSnapshot{value: val, isSet: true}
}

func (v NullableStoragePureVolumeSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureVolumeSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
