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

// checks if the StorageDiskSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageDiskSlot{}

// StorageDiskSlot Information of disk slots as reported by a storage controller.
type StorageDiskSlot struct {
	EquipmentSlot
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDisk         NullableStoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    NullableStorageControllerRelationship       `json:"StorageController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageDiskSlot StorageDiskSlot

// NewStorageDiskSlot instantiates a new StorageDiskSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskSlot(classId string, objectType string) *StorageDiskSlot {
	this := StorageDiskSlot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageDiskSlotWithDefaults instantiates a new StorageDiskSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskSlotWithDefaults() *StorageDiskSlot {
	this := StorageDiskSlot{}
	var classId string = "storage.DiskSlot"
	this.ClassId = classId
	var objectType string = "storage.DiskSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDiskSlot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDiskSlot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDiskSlot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.DiskSlot" of the ClassId field.
func (o *StorageDiskSlot) GetDefaultClassId() interface{} {
	return "storage.DiskSlot"
}

// GetObjectType returns the ObjectType field value
func (o *StorageDiskSlot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDiskSlot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDiskSlot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.DiskSlot" of the ObjectType field.
func (o *StorageDiskSlot) GetDefaultObjectType() interface{} {
	return "storage.DiskSlot"
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskSlot) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskSlot) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageDiskSlot) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageDiskSlot) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageDiskSlot) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageDiskSlot) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskSlot) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || IsNil(o.PhysicalDisk.Get()) {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk.Get()
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskSlot) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalDisk.Get(), o.PhysicalDisk.IsSet()
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StorageDiskSlot) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk.IsSet() {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given NullableStoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StorageDiskSlot) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk.Set(&v)
}

// SetPhysicalDiskNil sets the value for PhysicalDisk to be an explicit nil
func (o *StorageDiskSlot) SetPhysicalDiskNil() {
	o.PhysicalDisk.Set(nil)
}

// UnsetPhysicalDisk ensures that no value is present for PhysicalDisk, not even an explicit nil
func (o *StorageDiskSlot) UnsetPhysicalDisk() {
	o.PhysicalDisk.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageDiskSlot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageDiskSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageDiskSlot) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageDiskSlot) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageController returns the StorageController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskSlot) GetStorageController() StorageControllerRelationship {
	if o == nil || IsNil(o.StorageController.Get()) {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController.Get()
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskSlot) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageController.Get(), o.StorageController.IsSet()
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageDiskSlot) HasStorageController() bool {
	if o != nil && o.StorageController.IsSet() {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given NullableStorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageDiskSlot) SetStorageController(v StorageControllerRelationship) {
	o.StorageController.Set(&v)
}

// SetStorageControllerNil sets the value for StorageController to be an explicit nil
func (o *StorageDiskSlot) SetStorageControllerNil() {
	o.StorageController.Set(nil)
}

// UnsetStorageController ensures that no value is present for StorageController, not even an explicit nil
func (o *StorageDiskSlot) UnsetStorageController() {
	o.StorageController.Unset()
}

func (o StorageDiskSlot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageDiskSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentSlot, errEquipmentSlot := json.Marshal(o.EquipmentSlot)
	if errEquipmentSlot != nil {
		return map[string]interface{}{}, errEquipmentSlot
	}
	errEquipmentSlot = json.Unmarshal([]byte(serializedEquipmentSlot), &toSerialize)
	if errEquipmentSlot != nil {
		return map[string]interface{}{}, errEquipmentSlot
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.PhysicalDisk.IsSet() {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageController.IsSet() {
		toSerialize["StorageController"] = o.StorageController.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageDiskSlot) UnmarshalJSON(data []byte) (err error) {
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
	type StorageDiskSlotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                                      `json:"ObjectType"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		PhysicalDisk        NullableStoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StorageController   NullableStorageControllerRelationship       `json:"StorageController,omitempty"`
	}

	varStorageDiskSlotWithoutEmbeddedStruct := StorageDiskSlotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageDiskSlotWithoutEmbeddedStruct)
	if err == nil {
		varStorageDiskSlot := _StorageDiskSlot{}
		varStorageDiskSlot.ClassId = varStorageDiskSlotWithoutEmbeddedStruct.ClassId
		varStorageDiskSlot.ObjectType = varStorageDiskSlotWithoutEmbeddedStruct.ObjectType
		varStorageDiskSlot.InventoryDeviceInfo = varStorageDiskSlotWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageDiskSlot.PhysicalDisk = varStorageDiskSlotWithoutEmbeddedStruct.PhysicalDisk
		varStorageDiskSlot.RegisteredDevice = varStorageDiskSlotWithoutEmbeddedStruct.RegisteredDevice
		varStorageDiskSlot.StorageController = varStorageDiskSlotWithoutEmbeddedStruct.StorageController
		*o = StorageDiskSlot(varStorageDiskSlot)
	} else {
		return err
	}

	varStorageDiskSlot := _StorageDiskSlot{}

	err = json.Unmarshal(data, &varStorageDiskSlot)
	if err == nil {
		o.EquipmentSlot = varStorageDiskSlot.EquipmentSlot
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisk")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")

		// remove fields from embedded structs
		reflectEquipmentSlot := reflect.ValueOf(o.EquipmentSlot)
		for i := 0; i < reflectEquipmentSlot.Type().NumField(); i++ {
			t := reflectEquipmentSlot.Type().Field(i)

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

type NullableStorageDiskSlot struct {
	value *StorageDiskSlot
	isSet bool
}

func (v NullableStorageDiskSlot) Get() *StorageDiskSlot {
	return v.value
}

func (v *NullableStorageDiskSlot) Set(val *StorageDiskSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskSlot(val *StorageDiskSlot) *NullableStorageDiskSlot {
	return &NullableStorageDiskSlot{value: val, isSet: true}
}

func (v NullableStorageDiskSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
