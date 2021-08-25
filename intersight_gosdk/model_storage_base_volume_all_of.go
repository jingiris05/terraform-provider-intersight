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
)

// StorageBaseVolumeAllOf Definition of the list of properties defined in 'storage.BaseVolume', excluding properties defined in parent classes.
type StorageBaseVolumeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Short description about the volume.
	Description *string `json:"Description,omitempty"`
	// NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor.
	NaaId *string `json:"NaaId,omitempty"`
	// Named entity of the volume.
	Name *string `json:"Name,omitempty"`
	// User provisioned volume size. It is the size exposed to host.
	Size                 *int64                      `json:"Size,omitempty"`
	StorageUtilization   NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseVolumeAllOf StorageBaseVolumeAllOf

// NewStorageBaseVolumeAllOf instantiates a new StorageBaseVolumeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseVolumeAllOf(classId string, objectType string) *StorageBaseVolumeAllOf {
	this := StorageBaseVolumeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseVolumeAllOfWithDefaults instantiates a new StorageBaseVolumeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseVolumeAllOfWithDefaults() *StorageBaseVolumeAllOf {
	this := StorageBaseVolumeAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseVolumeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseVolumeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseVolumeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseVolumeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageBaseVolumeAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetNaaId returns the NaaId field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetNaaId() string {
	if o == nil || o.NaaId == nil {
		var ret string
		return ret
	}
	return *o.NaaId
}

// GetNaaIdOk returns a tuple with the NaaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetNaaIdOk() (*string, bool) {
	if o == nil || o.NaaId == nil {
		return nil, false
	}
	return o.NaaId, true
}

// HasNaaId returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasNaaId() bool {
	if o != nil && o.NaaId != nil {
		return true
	}

	return false
}

// SetNaaId gets a reference to the given string and assigns it to the NaaId field.
func (o *StorageBaseVolumeAllOf) SetNaaId(v string) {
	o.NaaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseVolumeAllOf) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseVolumeAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseVolumeAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseVolumeAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseVolumeAllOf) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseVolumeAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseVolumeAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseVolumeAllOf) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseVolumeAllOf) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseVolumeAllOf) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

func (o StorageBaseVolumeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.NaaId != nil {
		toSerialize["NaaId"] = o.NaaId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseVolumeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseVolumeAllOf := _StorageBaseVolumeAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseVolumeAllOf); err == nil {
		*o = StorageBaseVolumeAllOf(varStorageBaseVolumeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "NaaId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "StorageUtilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseVolumeAllOf struct {
	value *StorageBaseVolumeAllOf
	isSet bool
}

func (v NullableStorageBaseVolumeAllOf) Get() *StorageBaseVolumeAllOf {
	return v.value
}

func (v *NullableStorageBaseVolumeAllOf) Set(val *StorageBaseVolumeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseVolumeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseVolumeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseVolumeAllOf(val *StorageBaseVolumeAllOf) *NullableStorageBaseVolumeAllOf {
	return &NullableStorageBaseVolumeAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseVolumeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseVolumeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
