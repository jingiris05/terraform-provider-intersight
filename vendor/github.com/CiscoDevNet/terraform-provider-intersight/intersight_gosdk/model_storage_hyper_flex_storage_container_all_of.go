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
	"time"
)

// StorageHyperFlexStorageContainerAllOf Definition of the list of properties defined in 'storage.HyperFlexStorageContainer', excluding properties defined in parent classes.
type StorageHyperFlexStorageContainerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Storage Container data block size
	DataBlockSize *int64 `json:"DataBlockSize,omitempty"`
	// Indicates whether the Storage Container has Volumes.
	InUse *bool `json:"InUse,omitempty"`
	// Storage container's last access time.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// Storage container's last modified time.
	LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
	// Provisioned Capacity of the Storage container.
	ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
	// Provisioned Capacity Utilization of All Volumes associated with the Storage Container.
	ProvisionedVolumeCapacityUtilization *float32 `json:"ProvisionedVolumeCapacityUtilization,omitempty"`
	// Storage Container type (SMB/NFS/iSCSI). * `NFS` - Storage container created/accesed through NFS protocol. * `SMB` - Storage container created/accessed through SMB protocol. * `iSCSI` - Storage container created/accessed through iSCSI protocol.
	Type *string `json:"Type,omitempty"`
	// Uncompressed bytes on Storage Container.
	UnCompressedUsedBytes *int64 `json:"UnCompressedUsedBytes,omitempty"`
	// UUID of the Datastore/Storage Containter.
	Uuid *string `json:"Uuid,omitempty"`
	// Number of Volumes associated with the Storage Container.
	VolumeCount          *int64                               `json:"VolumeCount,omitempty"`
	Cluster              *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHyperFlexStorageContainerAllOf StorageHyperFlexStorageContainerAllOf

// NewStorageHyperFlexStorageContainerAllOf instantiates a new StorageHyperFlexStorageContainerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHyperFlexStorageContainerAllOf(classId string, objectType string) *StorageHyperFlexStorageContainerAllOf {
	this := StorageHyperFlexStorageContainerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHyperFlexStorageContainerAllOfWithDefaults instantiates a new StorageHyperFlexStorageContainerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHyperFlexStorageContainerAllOfWithDefaults() *StorageHyperFlexStorageContainerAllOf {
	this := StorageHyperFlexStorageContainerAllOf{}
	var classId string = "storage.HyperFlexStorageContainer"
	this.ClassId = classId
	var objectType string = "storage.HyperFlexStorageContainer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHyperFlexStorageContainerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHyperFlexStorageContainerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHyperFlexStorageContainerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHyperFlexStorageContainerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataBlockSize returns the DataBlockSize field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetDataBlockSize() int64 {
	if o == nil || o.DataBlockSize == nil {
		var ret int64
		return ret
	}
	return *o.DataBlockSize
}

// GetDataBlockSizeOk returns a tuple with the DataBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetDataBlockSizeOk() (*int64, bool) {
	if o == nil || o.DataBlockSize == nil {
		return nil, false
	}
	return o.DataBlockSize, true
}

// HasDataBlockSize returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasDataBlockSize() bool {
	if o != nil && o.DataBlockSize != nil {
		return true
	}

	return false
}

// SetDataBlockSize gets a reference to the given int64 and assigns it to the DataBlockSize field.
func (o *StorageHyperFlexStorageContainerAllOf) SetDataBlockSize(v int64) {
	o.DataBlockSize = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetInUse() bool {
	if o == nil || o.InUse == nil {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetInUseOk() (*bool, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *StorageHyperFlexStorageContainerAllOf) SetInUse(v bool) {
	o.InUse = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetLastAccessTime() time.Time {
	if o == nil || o.LastAccessTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *StorageHyperFlexStorageContainerAllOf) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *StorageHyperFlexStorageContainerAllOf) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedCapacity() int64 {
	if o == nil || o.ProvisionedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedCapacity
}

// GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedCapacityOk() (*int64, bool) {
	if o == nil || o.ProvisionedCapacity == nil {
		return nil, false
	}
	return o.ProvisionedCapacity, true
}

// HasProvisionedCapacity returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasProvisionedCapacity() bool {
	if o != nil && o.ProvisionedCapacity != nil {
		return true
	}

	return false
}

// SetProvisionedCapacity gets a reference to the given int64 and assigns it to the ProvisionedCapacity field.
func (o *StorageHyperFlexStorageContainerAllOf) SetProvisionedCapacity(v int64) {
	o.ProvisionedCapacity = &v
}

// GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilization() float32 {
	if o == nil || o.ProvisionedVolumeCapacityUtilization == nil {
		var ret float32
		return ret
	}
	return *o.ProvisionedVolumeCapacityUtilization
}

// GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool) {
	if o == nil || o.ProvisionedVolumeCapacityUtilization == nil {
		return nil, false
	}
	return o.ProvisionedVolumeCapacityUtilization, true
}

// HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasProvisionedVolumeCapacityUtilization() bool {
	if o != nil && o.ProvisionedVolumeCapacityUtilization != nil {
		return true
	}

	return false
}

// SetProvisionedVolumeCapacityUtilization gets a reference to the given float32 and assigns it to the ProvisionedVolumeCapacityUtilization field.
func (o *StorageHyperFlexStorageContainerAllOf) SetProvisionedVolumeCapacityUtilization(v float32) {
	o.ProvisionedVolumeCapacityUtilization = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageHyperFlexStorageContainerAllOf) SetType(v string) {
	o.Type = &v
}

// GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetUnCompressedUsedBytes() int64 {
	if o == nil || o.UnCompressedUsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UnCompressedUsedBytes
}

// GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetUnCompressedUsedBytesOk() (*int64, bool) {
	if o == nil || o.UnCompressedUsedBytes == nil {
		return nil, false
	}
	return o.UnCompressedUsedBytes, true
}

// HasUnCompressedUsedBytes returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasUnCompressedUsedBytes() bool {
	if o != nil && o.UnCompressedUsedBytes != nil {
		return true
	}

	return false
}

// SetUnCompressedUsedBytes gets a reference to the given int64 and assigns it to the UnCompressedUsedBytes field.
func (o *StorageHyperFlexStorageContainerAllOf) SetUnCompressedUsedBytes(v int64) {
	o.UnCompressedUsedBytes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageHyperFlexStorageContainerAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeCount returns the VolumeCount field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetVolumeCount() int64 {
	if o == nil || o.VolumeCount == nil {
		var ret int64
		return ret
	}
	return *o.VolumeCount
}

// GetVolumeCountOk returns a tuple with the VolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetVolumeCountOk() (*int64, bool) {
	if o == nil || o.VolumeCount == nil {
		return nil, false
	}
	return o.VolumeCount, true
}

// HasVolumeCount returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasVolumeCount() bool {
	if o != nil && o.VolumeCount != nil {
		return true
	}

	return false
}

// SetVolumeCount gets a reference to the given int64 and assigns it to the VolumeCount field.
func (o *StorageHyperFlexStorageContainerAllOf) SetVolumeCount(v int64) {
	o.VolumeCount = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *StorageHyperFlexStorageContainerAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHyperFlexStorageContainerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHyperFlexStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataBlockSize != nil {
		toSerialize["DataBlockSize"] = o.DataBlockSize
	}
	if o.InUse != nil {
		toSerialize["InUse"] = o.InUse
	}
	if o.LastAccessTime != nil {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if o.LastModifiedTime != nil {
		toSerialize["LastModifiedTime"] = o.LastModifiedTime
	}
	if o.ProvisionedCapacity != nil {
		toSerialize["ProvisionedCapacity"] = o.ProvisionedCapacity
	}
	if o.ProvisionedVolumeCapacityUtilization != nil {
		toSerialize["ProvisionedVolumeCapacityUtilization"] = o.ProvisionedVolumeCapacityUtilization
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.UnCompressedUsedBytes != nil {
		toSerialize["UnCompressedUsedBytes"] = o.UnCompressedUsedBytes
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeCount != nil {
		toSerialize["VolumeCount"] = o.VolumeCount
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHyperFlexStorageContainerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHyperFlexStorageContainerAllOf := _StorageHyperFlexStorageContainerAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHyperFlexStorageContainerAllOf); err == nil {
		*o = StorageHyperFlexStorageContainerAllOf(varStorageHyperFlexStorageContainerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataBlockSize")
		delete(additionalProperties, "InUse")
		delete(additionalProperties, "LastAccessTime")
		delete(additionalProperties, "LastModifiedTime")
		delete(additionalProperties, "ProvisionedCapacity")
		delete(additionalProperties, "ProvisionedVolumeCapacityUtilization")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "UnCompressedUsedBytes")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeCount")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHyperFlexStorageContainerAllOf struct {
	value *StorageHyperFlexStorageContainerAllOf
	isSet bool
}

func (v NullableStorageHyperFlexStorageContainerAllOf) Get() *StorageHyperFlexStorageContainerAllOf {
	return v.value
}

func (v *NullableStorageHyperFlexStorageContainerAllOf) Set(val *StorageHyperFlexStorageContainerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHyperFlexStorageContainerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHyperFlexStorageContainerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHyperFlexStorageContainerAllOf(val *StorageHyperFlexStorageContainerAllOf) *NullableStorageHyperFlexStorageContainerAllOf {
	return &NullableStorageHyperFlexStorageContainerAllOf{value: val, isSet: true}
}

func (v NullableStorageHyperFlexStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHyperFlexStorageContainerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
