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

// StorageNetAppLunAllOf Definition of the list of properties defined in 'storage.NetAppLun', excluding properties defined in parent classes.
type StorageNetAppLunAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                  `json:"ObjectType"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// The state of the volume and aggregate that contain the LUN. LUNs are only available when their containers are available.
	ContainerState *string `json:"ContainerState,omitempty"`
	// Reports if the LUN is mapped to one or more initiator groups.
	IsMapped *string `json:"IsMapped,omitempty"`
	// Unique identifier of LUN across data center.
	Key *string `json:"Key,omitempty"`
	// Reports if the LUN is mapped to one or more initiator groups.
	// Deprecated
	Mapped *bool `json:"Mapped,omitempty"`
	// The operating system (OS) type for this LUN. * `Linux` - Family of open source Unix-like operating systems based on the Linux kernel. * `AIX` - Advanced Interactive Executive (AIX). * `HP-UX` - HP-UX is implementation of the Unix operating system, based on Unix System V. * `Hyper-V` - Windows Server 2008 or Windows Server 2012 Hyper-V. * `OpenVMS` - OpenVMS is multi-user, multiprocessing virtual memory-based operating system. * `Solaris` - Solaris is a Unix operating system. * `NetWare` - NetWare is a computer network operating system. * `VMware` - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers. * `Windows` - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style. * `Xen` - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently.
	OsType *string `json:"OsType,omitempty"`
	// Path where the LUN is mounted.
	Path *string `json:"Path,omitempty"`
	// Serial number for the provisioned LUN.
	Serial *string `json:"Serial,omitempty"`
	// The administrative state of a LUN. * `offline` - The LUN is administratively offline, or a more detailed offline reason is not available. * `online` - The state of the LUN is online.
	State *string `json:"State,omitempty"`
	// The storage virtual machine name for the lun.
	SvmName *string `json:"SvmName,omitempty"`
	// Universally unique identifier of the LUN.
	Uuid *string `json:"Uuid,omitempty"`
	// The parent volume name for the lun.
	VolumeName *string                           `json:"VolumeName,omitempty"`
	Array      *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppLunEvent resources.
	Events []StorageNetAppLunEventRelationship `json:"Events,omitempty"`
	// An array of relationships to storageNetAppInitiatorGroup resources.
	Host                 []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
	StorageContainer     *StorageNetAppVolumeRelationship          `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLunAllOf StorageNetAppLunAllOf

// NewStorageNetAppLunAllOf instantiates a new StorageNetAppLunAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLunAllOf(classId string, objectType string) *StorageNetAppLunAllOf {
	this := StorageNetAppLunAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunAllOfWithDefaults instantiates a new StorageNetAppLunAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunAllOfWithDefaults() *StorageNetAppLunAllOf {
	this := StorageNetAppLunAllOf{}
	var classId string = "storage.NetAppLun"
	this.ClassId = classId
	var objectType string = "storage.NetAppLun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLunAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLunAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLunAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLunAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppLunAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetContainerState returns the ContainerState field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetContainerState() string {
	if o == nil || o.ContainerState == nil {
		var ret string
		return ret
	}
	return *o.ContainerState
}

// GetContainerStateOk returns a tuple with the ContainerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetContainerStateOk() (*string, bool) {
	if o == nil || o.ContainerState == nil {
		return nil, false
	}
	return o.ContainerState, true
}

// HasContainerState returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasContainerState() bool {
	if o != nil && o.ContainerState != nil {
		return true
	}

	return false
}

// SetContainerState gets a reference to the given string and assigns it to the ContainerState field.
func (o *StorageNetAppLunAllOf) SetContainerState(v string) {
	o.ContainerState = &v
}

// GetIsMapped returns the IsMapped field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetIsMapped() string {
	if o == nil || o.IsMapped == nil {
		var ret string
		return ret
	}
	return *o.IsMapped
}

// GetIsMappedOk returns a tuple with the IsMapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetIsMappedOk() (*string, bool) {
	if o == nil || o.IsMapped == nil {
		return nil, false
	}
	return o.IsMapped, true
}

// HasIsMapped returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasIsMapped() bool {
	if o != nil && o.IsMapped != nil {
		return true
	}

	return false
}

// SetIsMapped gets a reference to the given string and assigns it to the IsMapped field.
func (o *StorageNetAppLunAllOf) SetIsMapped(v string) {
	o.IsMapped = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppLunAllOf) SetKey(v string) {
	o.Key = &v
}

// GetMapped returns the Mapped field value if set, zero value otherwise.
// Deprecated
func (o *StorageNetAppLunAllOf) GetMapped() bool {
	if o == nil || o.Mapped == nil {
		var ret bool
		return ret
	}
	return *o.Mapped
}

// GetMappedOk returns a tuple with the Mapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StorageNetAppLunAllOf) GetMappedOk() (*bool, bool) {
	if o == nil || o.Mapped == nil {
		return nil, false
	}
	return o.Mapped, true
}

// HasMapped returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasMapped() bool {
	if o != nil && o.Mapped != nil {
		return true
	}

	return false
}

// SetMapped gets a reference to the given bool and assigns it to the Mapped field.
// Deprecated
func (o *StorageNetAppLunAllOf) SetMapped(v bool) {
	o.Mapped = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetOsType() string {
	if o == nil || o.OsType == nil {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetOsTypeOk() (*string, bool) {
	if o == nil || o.OsType == nil {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasOsType() bool {
	if o != nil && o.OsType != nil {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *StorageNetAppLunAllOf) SetOsType(v string) {
	o.OsType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StorageNetAppLunAllOf) SetPath(v string) {
	o.Path = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *StorageNetAppLunAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppLunAllOf) SetState(v string) {
	o.State = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetSvmName() string {
	if o == nil || o.SvmName == nil {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetSvmNameOk() (*string, bool) {
	if o == nil || o.SvmName == nil {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasSvmName() bool {
	if o != nil && o.SvmName != nil {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppLunAllOf) SetSvmName(v string) {
	o.SvmName = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppLunAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *StorageNetAppLunAllOf) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppLunAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLunAllOf) GetEvents() []StorageNetAppLunEventRelationship {
	if o == nil {
		var ret []StorageNetAppLunEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLunAllOf) GetEventsOk() ([]StorageNetAppLunEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppLunEventRelationship and assigns it to the Events field.
func (o *StorageNetAppLunAllOf) SetEvents(v []StorageNetAppLunEventRelationship) {
	o.Events = v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLunAllOf) GetHost() []StorageNetAppInitiatorGroupRelationship {
	if o == nil {
		var ret []StorageNetAppInitiatorGroupRelationship
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLunAllOf) GetHostOk() ([]StorageNetAppInitiatorGroupRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given []StorageNetAppInitiatorGroupRelationship and assigns it to the Host field.
func (o *StorageNetAppLunAllOf) SetHost(v []StorageNetAppInitiatorGroupRelationship) {
	o.Host = v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppLunAllOf) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppLunAllOf) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppLunAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

func (o StorageNetAppLunAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.ContainerState != nil {
		toSerialize["ContainerState"] = o.ContainerState
	}
	if o.IsMapped != nil {
		toSerialize["IsMapped"] = o.IsMapped
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Mapped != nil {
		toSerialize["Mapped"] = o.Mapped
	}
	if o.OsType != nil {
		toSerialize["OsType"] = o.OsType
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SvmName != nil {
		toSerialize["SvmName"] = o.SvmName
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeName != nil {
		toSerialize["VolumeName"] = o.VolumeName
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppLunAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppLunAllOf := _StorageNetAppLunAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppLunAllOf); err == nil {
		*o = StorageNetAppLunAllOf(varStorageNetAppLunAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "ContainerState")
		delete(additionalProperties, "IsMapped")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Mapped")
		delete(additionalProperties, "OsType")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "State")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeName")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "StorageContainer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppLunAllOf struct {
	value *StorageNetAppLunAllOf
	isSet bool
}

func (v NullableStorageNetAppLunAllOf) Get() *StorageNetAppLunAllOf {
	return v.value
}

func (v *NullableStorageNetAppLunAllOf) Set(val *StorageNetAppLunAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLunAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLunAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLunAllOf(val *StorageNetAppLunAllOf) *NullableStorageNetAppLunAllOf {
	return &NullableStorageNetAppLunAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppLunAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLunAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
