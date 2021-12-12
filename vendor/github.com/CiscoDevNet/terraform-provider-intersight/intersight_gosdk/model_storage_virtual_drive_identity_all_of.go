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
)

// StorageVirtualDriveIdentityAllOf Definition of the list of properties defined in 'storage.VirtualDriveIdentity', excluding properties defined in parent classes.
type StorageVirtualDriveIdentityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The VirtualDrive Name which belongs to the Storage VirtualDrive.
	Name                 *string                           `json:"Name,omitempty"`
	ServerProfile        *ServerProfileRelationship        `json:"ServerProfile,omitempty"`
	StoragePolicy        *StorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
	VirtualDrive         *StorageVirtualDriveRelationship  `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveIdentityAllOf StorageVirtualDriveIdentityAllOf

// NewStorageVirtualDriveIdentityAllOf instantiates a new StorageVirtualDriveIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveIdentityAllOf(classId string, objectType string) *StorageVirtualDriveIdentityAllOf {
	this := StorageVirtualDriveIdentityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveIdentityAllOfWithDefaults instantiates a new StorageVirtualDriveIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveIdentityAllOfWithDefaults() *StorageVirtualDriveIdentityAllOf {
	this := StorageVirtualDriveIdentityAllOf{}
	var classId string = "storage.VirtualDriveIdentity"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveIdentityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveIdentityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveIdentityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveIdentityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentityAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentityAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveIdentityAllOf) SetName(v string) {
	o.Name = &v
}

// GetServerProfile returns the ServerProfile field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentityAllOf) GetServerProfile() ServerProfileRelationship {
	if o == nil || o.ServerProfile == nil {
		var ret ServerProfileRelationship
		return ret
	}
	return *o.ServerProfile
}

// GetServerProfileOk returns a tuple with the ServerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetServerProfileOk() (*ServerProfileRelationship, bool) {
	if o == nil || o.ServerProfile == nil {
		return nil, false
	}
	return o.ServerProfile, true
}

// HasServerProfile returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentityAllOf) HasServerProfile() bool {
	if o != nil && o.ServerProfile != nil {
		return true
	}

	return false
}

// SetServerProfile gets a reference to the given ServerProfileRelationship and assigns it to the ServerProfile field.
func (o *StorageVirtualDriveIdentityAllOf) SetServerProfile(v ServerProfileRelationship) {
	o.ServerProfile = &v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentityAllOf) GetStoragePolicy() StorageStoragePolicyRelationship {
	if o == nil || o.StoragePolicy == nil {
		var ret StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentityAllOf) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given StorageStoragePolicyRelationship and assigns it to the StoragePolicy field.
func (o *StorageVirtualDriveIdentityAllOf) SetStoragePolicy(v StorageStoragePolicyRelationship) {
	o.StoragePolicy = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentityAllOf) GetVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || o.VirtualDrive == nil {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentityAllOf) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentityAllOf) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveIdentityAllOf) SetVirtualDrive(v StorageVirtualDriveRelationship) {
	o.VirtualDrive = &v
}

func (o StorageVirtualDriveIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServerProfile != nil {
		toSerialize["ServerProfile"] = o.ServerProfile
	}
	if o.StoragePolicy != nil {
		toSerialize["StoragePolicy"] = o.StoragePolicy
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageVirtualDriveIdentityAllOf := _StorageVirtualDriveIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varStorageVirtualDriveIdentityAllOf); err == nil {
		*o = StorageVirtualDriveIdentityAllOf(varStorageVirtualDriveIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServerProfile")
		delete(additionalProperties, "StoragePolicy")
		delete(additionalProperties, "VirtualDrive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageVirtualDriveIdentityAllOf struct {
	value *StorageVirtualDriveIdentityAllOf
	isSet bool
}

func (v NullableStorageVirtualDriveIdentityAllOf) Get() *StorageVirtualDriveIdentityAllOf {
	return v.value
}

func (v *NullableStorageVirtualDriveIdentityAllOf) Set(val *StorageVirtualDriveIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveIdentityAllOf(val *StorageVirtualDriveIdentityAllOf) *NullableStorageVirtualDriveIdentityAllOf {
	return &NullableStorageVirtualDriveIdentityAllOf{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
