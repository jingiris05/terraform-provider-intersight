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

// StoragePureDisk Disk entity associated with Pure FlashArray.
type StoragePureDisk struct {
	StorageBaseArrayDisk
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Array                *StoragePureArrayRelationship        `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureDisk StoragePureDisk

// NewStoragePureDisk instantiates a new StoragePureDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureDisk(classId string, objectType string) *StoragePureDisk {
	this := StoragePureDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureDiskWithDefaults instantiates a new StoragePureDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureDiskWithDefaults() *StoragePureDisk {
	this := StoragePureDisk{}
	var classId string = "storage.PureDisk"
	this.ClassId = classId
	var objectType string = "storage.PureDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureDisk) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureDisk) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureDisk) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureDisk) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureDisk) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureDisk) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureDisk) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureDisk) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArrayDisk, errStorageBaseArrayDisk := json.Marshal(o.StorageBaseArrayDisk)
	if errStorageBaseArrayDisk != nil {
		return []byte{}, errStorageBaseArrayDisk
	}
	errStorageBaseArrayDisk = json.Unmarshal([]byte(serializedStorageBaseArrayDisk), &toSerialize)
	if errStorageBaseArrayDisk != nil {
		return []byte{}, errStorageBaseArrayDisk
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureDisk) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                               `json:"ObjectType"`
		Array            *StoragePureArrayRelationship        `json:"Array,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureDiskWithoutEmbeddedStruct := StoragePureDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureDiskWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureDisk := _StoragePureDisk{}
		varStoragePureDisk.ClassId = varStoragePureDiskWithoutEmbeddedStruct.ClassId
		varStoragePureDisk.ObjectType = varStoragePureDiskWithoutEmbeddedStruct.ObjectType
		varStoragePureDisk.Array = varStoragePureDiskWithoutEmbeddedStruct.Array
		varStoragePureDisk.RegisteredDevice = varStoragePureDiskWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureDisk(varStoragePureDisk)
	} else {
		return err
	}

	varStoragePureDisk := _StoragePureDisk{}

	err = json.Unmarshal(bytes, &varStoragePureDisk)
	if err == nil {
		o.StorageBaseArrayDisk = varStoragePureDisk.StorageBaseArrayDisk
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseArrayDisk := reflect.ValueOf(o.StorageBaseArrayDisk)
		for i := 0; i < reflectStorageBaseArrayDisk.Type().NumField(); i++ {
			t := reflectStorageBaseArrayDisk.Type().Field(i)

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

type NullableStoragePureDisk struct {
	value *StoragePureDisk
	isSet bool
}

func (v NullableStoragePureDisk) Get() *StoragePureDisk {
	return v.value
}

func (v *NullableStoragePureDisk) Set(val *StoragePureDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureDisk(val *StoragePureDisk) *NullableStoragePureDisk {
	return &NullableStoragePureDisk{value: val, isSet: true}
}

func (v NullableStoragePureDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
