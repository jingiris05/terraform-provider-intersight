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

// checks if the StorageNetAppNamespace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppNamespace{}

// StorageNetAppNamespace NetApp Namespace is a collection of addressable logical blocks presented to hosts connected to the storage virtual machine using the NVMe over Fabrics protocol.
type StorageNetAppNamespace struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available.
	ContainerState *string `json:"ContainerState,omitempty"`
	// Reports if the NVMe namespace is mapped to an NVMe subsystem.
	IsMapped *string `json:"IsMapped,omitempty"`
	// The base name component of the NVMe namespace.
	Name *string `json:"Name,omitempty"`
	// The NVMe namespace identifier. An identifier used by an NVMe controller to provide access to the NVMe namespace. The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \"h\".
	NamespaceId *string `json:"NamespaceId,omitempty"`
	// The fully qualified path name of the NVMe namespace composed of a \"/vol\" prefix, the volume name, the (optional) qtree name and base name of the namespace.
	Path *string `json:"Path,omitempty"`
	// The state of the NVMe namespace. Normal states for a namespace are online and offline. Other states indicate errors (nvfail, space_error).
	State              *string                     `json:"State,omitempty"`
	StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	// The NVMe subsystem name to which the NVMe namespace is mapped.
	SubsystemName *string `json:"SubsystemName,omitempty"`
	// The storage virtual machine name for the NVMe namespace.
	SvmName *string `json:"SvmName,omitempty"`
	// Universally unique identifier of the NVMe namespace.
	Uuid *string `json:"Uuid,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// The volume name in which the NVMe namespace is located.
	VolumeName           *string                                  `json:"VolumeName,omitempty"`
	Array                NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	StorageContainer     NullableStorageNetAppVolumeRelationship  `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNamespace StorageNetAppNamespace

// NewStorageNetAppNamespace instantiates a new StorageNetAppNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNamespace(classId string, objectType string) *StorageNetAppNamespace {
	this := StorageNetAppNamespace{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNamespaceWithDefaults instantiates a new StorageNetAppNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNamespaceWithDefaults() *StorageNetAppNamespace {
	this := StorageNetAppNamespace{}
	var classId string = "storage.NetAppNamespace"
	this.ClassId = classId
	var objectType string = "storage.NetAppNamespace"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNamespace) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNamespace) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppNamespace" of the ClassId field.
func (o *StorageNetAppNamespace) GetDefaultClassId() interface{} {
	return "storage.NetAppNamespace"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNamespace) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNamespace) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppNamespace" of the ObjectType field.
func (o *StorageNetAppNamespace) GetDefaultObjectType() interface{} {
	return "storage.NetAppNamespace"
}

// GetContainerState returns the ContainerState field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetContainerState() string {
	if o == nil || IsNil(o.ContainerState) {
		var ret string
		return ret
	}
	return *o.ContainerState
}

// GetContainerStateOk returns a tuple with the ContainerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetContainerStateOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerState) {
		return nil, false
	}
	return o.ContainerState, true
}

// HasContainerState returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasContainerState() bool {
	if o != nil && !IsNil(o.ContainerState) {
		return true
	}

	return false
}

// SetContainerState gets a reference to the given string and assigns it to the ContainerState field.
func (o *StorageNetAppNamespace) SetContainerState(v string) {
	o.ContainerState = &v
}

// GetIsMapped returns the IsMapped field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetIsMapped() string {
	if o == nil || IsNil(o.IsMapped) {
		var ret string
		return ret
	}
	return *o.IsMapped
}

// GetIsMappedOk returns a tuple with the IsMapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetIsMappedOk() (*string, bool) {
	if o == nil || IsNil(o.IsMapped) {
		return nil, false
	}
	return o.IsMapped, true
}

// HasIsMapped returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasIsMapped() bool {
	if o != nil && !IsNil(o.IsMapped) {
		return true
	}

	return false
}

// SetIsMapped gets a reference to the given string and assigns it to the IsMapped field.
func (o *StorageNetAppNamespace) SetIsMapped(v string) {
	o.IsMapped = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppNamespace) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetNamespaceId() string {
	if o == nil || IsNil(o.NamespaceId) {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceId) {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasNamespaceId() bool {
	if o != nil && !IsNil(o.NamespaceId) {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *StorageNetAppNamespace) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StorageNetAppNamespace) SetPath(v string) {
	o.Path = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppNamespace) SetState(v string) {
	o.State = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNamespace) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || IsNil(o.StorageUtilization.Get()) {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNamespace) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageNetAppNamespace) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageNetAppNamespace) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageNetAppNamespace) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

// GetSubsystemName returns the SubsystemName field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetSubsystemName() string {
	if o == nil || IsNil(o.SubsystemName) {
		var ret string
		return ret
	}
	return *o.SubsystemName
}

// GetSubsystemNameOk returns a tuple with the SubsystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetSubsystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubsystemName) {
		return nil, false
	}
	return o.SubsystemName, true
}

// HasSubsystemName returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasSubsystemName() bool {
	if o != nil && !IsNil(o.SubsystemName) {
		return true
	}

	return false
}

// SetSubsystemName gets a reference to the given string and assigns it to the SubsystemName field.
func (o *StorageNetAppNamespace) SetSubsystemName(v string) {
	o.SubsystemName = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetSvmName() string {
	if o == nil || IsNil(o.SvmName) {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetSvmNameOk() (*string, bool) {
	if o == nil || IsNil(o.SvmName) {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasSvmName() bool {
	if o != nil && !IsNil(o.SvmName) {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppNamespace) SetSvmName(v string) {
	o.SvmName = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppNamespace) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *StorageNetAppNamespace) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNamespace) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *StorageNetAppNamespace) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNamespace) GetArray() StorageNetAppClusterRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNamespace) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNamespace) SetArray(v StorageNetAppClusterRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageNetAppNamespace) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageNetAppNamespace) UnsetArray() {
	o.Array.Unset()
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNamespace) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || IsNil(o.StorageContainer.Get()) {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer.Get()
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNamespace) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageContainer.Get(), o.StorageContainer.IsSet()
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppNamespace) HasStorageContainer() bool {
	if o != nil && o.StorageContainer.IsSet() {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given NullableStorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppNamespace) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer.Set(&v)
}

// SetStorageContainerNil sets the value for StorageContainer to be an explicit nil
func (o *StorageNetAppNamespace) SetStorageContainerNil() {
	o.StorageContainer.Set(nil)
}

// UnsetStorageContainer ensures that no value is present for StorageContainer, not even an explicit nil
func (o *StorageNetAppNamespace) UnsetStorageContainer() {
	o.StorageContainer.Unset()
}

func (o StorageNetAppNamespace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppNamespace) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ContainerState) {
		toSerialize["ContainerState"] = o.ContainerState
	}
	if !IsNil(o.IsMapped) {
		toSerialize["IsMapped"] = o.IsMapped
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NamespaceId) {
		toSerialize["NamespaceId"] = o.NamespaceId
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}
	if !IsNil(o.SubsystemName) {
		toSerialize["SubsystemName"] = o.SubsystemName
	}
	if !IsNil(o.SvmName) {
		toSerialize["SvmName"] = o.SvmName
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.VolumeName) {
		toSerialize["VolumeName"] = o.VolumeName
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.StorageContainer.IsSet() {
		toSerialize["StorageContainer"] = o.StorageContainer.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppNamespace) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppNamespaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available.
		ContainerState *string `json:"ContainerState,omitempty"`
		// Reports if the NVMe namespace is mapped to an NVMe subsystem.
		IsMapped *string `json:"IsMapped,omitempty"`
		// The base name component of the NVMe namespace.
		Name *string `json:"Name,omitempty"`
		// The NVMe namespace identifier. An identifier used by an NVMe controller to provide access to the NVMe namespace. The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \"h\".
		NamespaceId *string `json:"NamespaceId,omitempty"`
		// The fully qualified path name of the NVMe namespace composed of a \"/vol\" prefix, the volume name, the (optional) qtree name and base name of the namespace.
		Path *string `json:"Path,omitempty"`
		// The state of the NVMe namespace. Normal states for a namespace are online and offline. Other states indicate errors (nvfail, space_error).
		State              *string                     `json:"State,omitempty"`
		StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
		// The NVMe subsystem name to which the NVMe namespace is mapped.
		SubsystemName *string `json:"SubsystemName,omitempty"`
		// The storage virtual machine name for the NVMe namespace.
		SvmName *string `json:"SvmName,omitempty"`
		// Universally unique identifier of the NVMe namespace.
		Uuid *string `json:"Uuid,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		// The volume name in which the NVMe namespace is located.
		VolumeName       *string                                  `json:"VolumeName,omitempty"`
		Array            NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
		StorageContainer NullableStorageNetAppVolumeRelationship  `json:"StorageContainer,omitempty"`
	}

	varStorageNetAppNamespaceWithoutEmbeddedStruct := StorageNetAppNamespaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppNamespaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppNamespace := _StorageNetAppNamespace{}
		varStorageNetAppNamespace.ClassId = varStorageNetAppNamespaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppNamespace.ObjectType = varStorageNetAppNamespaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppNamespace.ContainerState = varStorageNetAppNamespaceWithoutEmbeddedStruct.ContainerState
		varStorageNetAppNamespace.IsMapped = varStorageNetAppNamespaceWithoutEmbeddedStruct.IsMapped
		varStorageNetAppNamespace.Name = varStorageNetAppNamespaceWithoutEmbeddedStruct.Name
		varStorageNetAppNamespace.NamespaceId = varStorageNetAppNamespaceWithoutEmbeddedStruct.NamespaceId
		varStorageNetAppNamespace.Path = varStorageNetAppNamespaceWithoutEmbeddedStruct.Path
		varStorageNetAppNamespace.State = varStorageNetAppNamespaceWithoutEmbeddedStruct.State
		varStorageNetAppNamespace.StorageUtilization = varStorageNetAppNamespaceWithoutEmbeddedStruct.StorageUtilization
		varStorageNetAppNamespace.SubsystemName = varStorageNetAppNamespaceWithoutEmbeddedStruct.SubsystemName
		varStorageNetAppNamespace.SvmName = varStorageNetAppNamespaceWithoutEmbeddedStruct.SvmName
		varStorageNetAppNamespace.Uuid = varStorageNetAppNamespaceWithoutEmbeddedStruct.Uuid
		varStorageNetAppNamespace.VolumeName = varStorageNetAppNamespaceWithoutEmbeddedStruct.VolumeName
		varStorageNetAppNamespace.Array = varStorageNetAppNamespaceWithoutEmbeddedStruct.Array
		varStorageNetAppNamespace.StorageContainer = varStorageNetAppNamespaceWithoutEmbeddedStruct.StorageContainer
		*o = StorageNetAppNamespace(varStorageNetAppNamespace)
	} else {
		return err
	}

	varStorageNetAppNamespace := _StorageNetAppNamespace{}

	err = json.Unmarshal(data, &varStorageNetAppNamespace)
	if err == nil {
		o.MoBaseMo = varStorageNetAppNamespace.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContainerState")
		delete(additionalProperties, "IsMapped")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NamespaceId")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "State")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "SubsystemName")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeName")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "StorageContainer")

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

type NullableStorageNetAppNamespace struct {
	value *StorageNetAppNamespace
	isSet bool
}

func (v NullableStorageNetAppNamespace) Get() *StorageNetAppNamespace {
	return v.value
}

func (v *NullableStorageNetAppNamespace) Set(val *StorageNetAppNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNamespace(val *StorageNetAppNamespace) *NullableStorageNetAppNamespace {
	return &NullableStorageNetAppNamespace{value: val, isSet: true}
}

func (v NullableStorageNetAppNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
