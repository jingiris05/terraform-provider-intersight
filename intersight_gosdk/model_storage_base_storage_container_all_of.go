/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// StorageBaseStorageContainerAllOf Definition of the list of properties defined in 'storage.BaseStorageContainer', excluding properties defined in parent classes.
type StorageBaseStorageContainerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Storage container's creation time.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Name of the storage container.
	Name                 *string                     `json:"Name,omitempty"`
	StorageUtilization   NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseStorageContainerAllOf StorageBaseStorageContainerAllOf

// NewStorageBaseStorageContainerAllOf instantiates a new StorageBaseStorageContainerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseStorageContainerAllOf(classId string, objectType string) *StorageBaseStorageContainerAllOf {
	this := StorageBaseStorageContainerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseStorageContainerAllOfWithDefaults instantiates a new StorageBaseStorageContainerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseStorageContainerAllOfWithDefaults() *StorageBaseStorageContainerAllOf {
	this := StorageBaseStorageContainerAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseStorageContainerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseStorageContainerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseStorageContainerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseStorageContainerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseStorageContainerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseStorageContainerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageBaseStorageContainerAllOf) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseStorageContainerAllOf) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageBaseStorageContainerAllOf) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageBaseStorageContainerAllOf) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseStorageContainerAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseStorageContainerAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseStorageContainerAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseStorageContainerAllOf) SetName(v string) {
	o.Name = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseStorageContainerAllOf) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseStorageContainerAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseStorageContainerAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseStorageContainerAllOf) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseStorageContainerAllOf) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseStorageContainerAllOf) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

func (o StorageBaseStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CreatedTime != nil {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseStorageContainerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseStorageContainerAllOf := _StorageBaseStorageContainerAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseStorageContainerAllOf); err == nil {
		*o = StorageBaseStorageContainerAllOf(varStorageBaseStorageContainerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StorageUtilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseStorageContainerAllOf struct {
	value *StorageBaseStorageContainerAllOf
	isSet bool
}

func (v NullableStorageBaseStorageContainerAllOf) Get() *StorageBaseStorageContainerAllOf {
	return v.value
}

func (v *NullableStorageBaseStorageContainerAllOf) Set(val *StorageBaseStorageContainerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseStorageContainerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseStorageContainerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseStorageContainerAllOf(val *StorageBaseStorageContainerAllOf) *NullableStorageBaseStorageContainerAllOf {
	return &NullableStorageBaseStorageContainerAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseStorageContainerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseStorageContainerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
