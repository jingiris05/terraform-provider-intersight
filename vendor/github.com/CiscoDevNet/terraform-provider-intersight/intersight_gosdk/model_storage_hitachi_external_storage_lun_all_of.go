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
)

// StorageHitachiExternalStorageLunAllOf Definition of the list of properties defined in 'storage.HitachiExternalStorageLun', excluding properties defined in parent classes.
type StorageHitachiExternalStorageLunAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// LUN that can be referenced from the port on the external storage system.
	ExternalLun *int64 `json:"ExternalLun,omitempty"`
	// Object ID of ports on an external storage system.
	ExternalPortId *string `json:"ExternalPortId,omitempty"`
	// Capacity of the external volume on the external storage system (1 block = 512 bytes).
	ExternalVolumeCapacity *int64 `json:"ExternalVolumeCapacity,omitempty"`
	// The product ID and the device identification (output in ASCII format) in the SCSI information for the external volume on the external storage system. This information is obtained in a format in which the product ID and the device identification are concatenated by a space.
	ExternalVolumeInfo *string `json:"ExternalVolumeInfo,omitempty"`
	// WWN of the external storage port.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// The iSCSI IP Address of the external storage port.
	IscsiIpAddress *string `json:"IscsiIpAddress,omitempty"`
	// The iSCSI Name of the external storage port.
	IscsiName *string `json:"IscsiName,omitempty"`
	// Port ID of the local storage.
	PortId *string `json:"PortId,omitempty"`
	// Virtual port ID. This attribute is displayed when an iSCSI port is used and virtual port mode is enabled.
	VirtualPortId        *int64                                         `json:"VirtualPortId,omitempty"`
	Array                *StorageHitachiArrayRelationship               `json:"Array,omitempty"`
	ExternalStoragePort  *StorageHitachiExternalStoragePortRelationship `json:"ExternalStoragePort,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship           `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiExternalStorageLunAllOf StorageHitachiExternalStorageLunAllOf

// NewStorageHitachiExternalStorageLunAllOf instantiates a new StorageHitachiExternalStorageLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiExternalStorageLunAllOf(classId string, objectType string) *StorageHitachiExternalStorageLunAllOf {
	this := StorageHitachiExternalStorageLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiExternalStorageLunAllOfWithDefaults instantiates a new StorageHitachiExternalStorageLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiExternalStorageLunAllOfWithDefaults() *StorageHitachiExternalStorageLunAllOf {
	this := StorageHitachiExternalStorageLunAllOf{}
	var classId string = "storage.HitachiExternalStorageLun"
	this.ClassId = classId
	var objectType string = "storage.HitachiExternalStorageLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiExternalStorageLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiExternalStorageLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiExternalStorageLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiExternalStorageLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalLun returns the ExternalLun field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalLun() int64 {
	if o == nil || o.ExternalLun == nil {
		var ret int64
		return ret
	}
	return *o.ExternalLun
}

// GetExternalLunOk returns a tuple with the ExternalLun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalLunOk() (*int64, bool) {
	if o == nil || o.ExternalLun == nil {
		return nil, false
	}
	return o.ExternalLun, true
}

// HasExternalLun returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalLun() bool {
	if o != nil && o.ExternalLun != nil {
		return true
	}

	return false
}

// SetExternalLun gets a reference to the given int64 and assigns it to the ExternalLun field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalLun(v int64) {
	o.ExternalLun = &v
}

// GetExternalPortId returns the ExternalPortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalPortId() string {
	if o == nil || o.ExternalPortId == nil {
		var ret string
		return ret
	}
	return *o.ExternalPortId
}

// GetExternalPortIdOk returns a tuple with the ExternalPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalPortIdOk() (*string, bool) {
	if o == nil || o.ExternalPortId == nil {
		return nil, false
	}
	return o.ExternalPortId, true
}

// HasExternalPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalPortId() bool {
	if o != nil && o.ExternalPortId != nil {
		return true
	}

	return false
}

// SetExternalPortId gets a reference to the given string and assigns it to the ExternalPortId field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalPortId(v string) {
	o.ExternalPortId = &v
}

// GetExternalVolumeCapacity returns the ExternalVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeCapacity() int64 {
	if o == nil || o.ExternalVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ExternalVolumeCapacity
}

// GetExternalVolumeCapacityOk returns a tuple with the ExternalVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.ExternalVolumeCapacity == nil {
		return nil, false
	}
	return o.ExternalVolumeCapacity, true
}

// HasExternalVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalVolumeCapacity() bool {
	if o != nil && o.ExternalVolumeCapacity != nil {
		return true
	}

	return false
}

// SetExternalVolumeCapacity gets a reference to the given int64 and assigns it to the ExternalVolumeCapacity field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalVolumeCapacity(v int64) {
	o.ExternalVolumeCapacity = &v
}

// GetExternalVolumeInfo returns the ExternalVolumeInfo field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeInfo() string {
	if o == nil || o.ExternalVolumeInfo == nil {
		var ret string
		return ret
	}
	return *o.ExternalVolumeInfo
}

// GetExternalVolumeInfoOk returns a tuple with the ExternalVolumeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalVolumeInfoOk() (*string, bool) {
	if o == nil || o.ExternalVolumeInfo == nil {
		return nil, false
	}
	return o.ExternalVolumeInfo, true
}

// HasExternalVolumeInfo returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalVolumeInfo() bool {
	if o != nil && o.ExternalVolumeInfo != nil {
		return true
	}

	return false
}

// SetExternalVolumeInfo gets a reference to the given string and assigns it to the ExternalVolumeInfo field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalVolumeInfo(v string) {
	o.ExternalVolumeInfo = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalWwn() string {
	if o == nil || o.ExternalWwn == nil {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalWwnOk() (*string, bool) {
	if o == nil || o.ExternalWwn == nil {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalWwn() bool {
	if o != nil && o.ExternalWwn != nil {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetIscsiIpAddress returns the IscsiIpAddress field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiIpAddress() string {
	if o == nil || o.IscsiIpAddress == nil {
		var ret string
		return ret
	}
	return *o.IscsiIpAddress
}

// GetIscsiIpAddressOk returns a tuple with the IscsiIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiIpAddressOk() (*string, bool) {
	if o == nil || o.IscsiIpAddress == nil {
		return nil, false
	}
	return o.IscsiIpAddress, true
}

// HasIscsiIpAddress returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasIscsiIpAddress() bool {
	if o != nil && o.IscsiIpAddress != nil {
		return true
	}

	return false
}

// SetIscsiIpAddress gets a reference to the given string and assigns it to the IscsiIpAddress field.
func (o *StorageHitachiExternalStorageLunAllOf) SetIscsiIpAddress(v string) {
	o.IscsiIpAddress = &v
}

// GetIscsiName returns the IscsiName field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiName() string {
	if o == nil || o.IscsiName == nil {
		var ret string
		return ret
	}
	return *o.IscsiName
}

// GetIscsiNameOk returns a tuple with the IscsiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetIscsiNameOk() (*string, bool) {
	if o == nil || o.IscsiName == nil {
		return nil, false
	}
	return o.IscsiName, true
}

// HasIscsiName returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasIscsiName() bool {
	if o != nil && o.IscsiName != nil {
		return true
	}

	return false
}

// SetIscsiName gets a reference to the given string and assigns it to the IscsiName field.
func (o *StorageHitachiExternalStorageLunAllOf) SetIscsiName(v string) {
	o.IscsiName = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageHitachiExternalStorageLunAllOf) SetPortId(v string) {
	o.PortId = &v
}

// GetVirtualPortId returns the VirtualPortId field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetVirtualPortId() int64 {
	if o == nil || o.VirtualPortId == nil {
		var ret int64
		return ret
	}
	return *o.VirtualPortId
}

// GetVirtualPortIdOk returns a tuple with the VirtualPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetVirtualPortIdOk() (*int64, bool) {
	if o == nil || o.VirtualPortId == nil {
		return nil, false
	}
	return o.VirtualPortId, true
}

// HasVirtualPortId returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasVirtualPortId() bool {
	if o != nil && o.VirtualPortId != nil {
		return true
	}

	return false
}

// SetVirtualPortId gets a reference to the given int64 and assigns it to the VirtualPortId field.
func (o *StorageHitachiExternalStorageLunAllOf) SetVirtualPortId(v int64) {
	o.VirtualPortId = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiExternalStorageLunAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetExternalStoragePort returns the ExternalStoragePort field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalStoragePort() StorageHitachiExternalStoragePortRelationship {
	if o == nil || o.ExternalStoragePort == nil {
		var ret StorageHitachiExternalStoragePortRelationship
		return ret
	}
	return *o.ExternalStoragePort
}

// GetExternalStoragePortOk returns a tuple with the ExternalStoragePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetExternalStoragePortOk() (*StorageHitachiExternalStoragePortRelationship, bool) {
	if o == nil || o.ExternalStoragePort == nil {
		return nil, false
	}
	return o.ExternalStoragePort, true
}

// HasExternalStoragePort returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasExternalStoragePort() bool {
	if o != nil && o.ExternalStoragePort != nil {
		return true
	}

	return false
}

// SetExternalStoragePort gets a reference to the given StorageHitachiExternalStoragePortRelationship and assigns it to the ExternalStoragePort field.
func (o *StorageHitachiExternalStorageLunAllOf) SetExternalStoragePort(v StorageHitachiExternalStoragePortRelationship) {
	o.ExternalStoragePort = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiExternalStorageLunAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalStorageLunAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiExternalStorageLunAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiExternalStorageLunAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiExternalStorageLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExternalLun != nil {
		toSerialize["ExternalLun"] = o.ExternalLun
	}
	if o.ExternalPortId != nil {
		toSerialize["ExternalPortId"] = o.ExternalPortId
	}
	if o.ExternalVolumeCapacity != nil {
		toSerialize["ExternalVolumeCapacity"] = o.ExternalVolumeCapacity
	}
	if o.ExternalVolumeInfo != nil {
		toSerialize["ExternalVolumeInfo"] = o.ExternalVolumeInfo
	}
	if o.ExternalWwn != nil {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if o.IscsiIpAddress != nil {
		toSerialize["IscsiIpAddress"] = o.IscsiIpAddress
	}
	if o.IscsiName != nil {
		toSerialize["IscsiName"] = o.IscsiName
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.VirtualPortId != nil {
		toSerialize["VirtualPortId"] = o.VirtualPortId
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ExternalStoragePort != nil {
		toSerialize["ExternalStoragePort"] = o.ExternalStoragePort
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiExternalStorageLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiExternalStorageLunAllOf := _StorageHitachiExternalStorageLunAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiExternalStorageLunAllOf); err == nil {
		*o = StorageHitachiExternalStorageLunAllOf(varStorageHitachiExternalStorageLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalLun")
		delete(additionalProperties, "ExternalPortId")
		delete(additionalProperties, "ExternalVolumeCapacity")
		delete(additionalProperties, "ExternalVolumeInfo")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "IscsiIpAddress")
		delete(additionalProperties, "IscsiName")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "VirtualPortId")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ExternalStoragePort")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiExternalStorageLunAllOf struct {
	value *StorageHitachiExternalStorageLunAllOf
	isSet bool
}

func (v NullableStorageHitachiExternalStorageLunAllOf) Get() *StorageHitachiExternalStorageLunAllOf {
	return v.value
}

func (v *NullableStorageHitachiExternalStorageLunAllOf) Set(val *StorageHitachiExternalStorageLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiExternalStorageLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiExternalStorageLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiExternalStorageLunAllOf(val *StorageHitachiExternalStorageLunAllOf) *NullableStorageHitachiExternalStorageLunAllOf {
	return &NullableStorageHitachiExternalStorageLunAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiExternalStorageLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiExternalStorageLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
